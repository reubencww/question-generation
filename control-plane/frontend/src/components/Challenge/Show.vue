<script setup lang="ts">
import { useRoute } from "vue-router";
import http from "@/http";
import { onBeforeMount, onBeforeUnmount, ref } from "vue";
import { Challenge } from "@/challenges";
import { PencilSquareIcon, TrashIcon } from "@heroicons/vue/24/solid";
import { Tippy } from "vue-tippy";
import { fuzzyDate } from "@/dates";
import Sockette from "sockette";

interface ChallengeResponse {
    data: Challenge;
}

const route = useRoute();

const challenge = ref<Challenge | null>(null);
const questionGenerationCorpus = ref("");

let ws: Sockette | null = null;

const fetchChallenge = async () => {
    let id = parseInt(<string>route.params.id);
    let response = await http.get<ChallengeResponse>(`/challenge/${id}`);
    challenge.value = response.data.data;
};

const updateQuestion = async (questionId: number, question: string, answer: string) => {
    await http.patch(`/challenge/question/${questionId}`, {
        question,
        answer,
    });
    await fetchChallenge();
};

const deleteQuestion = async (questionId: number) => {
    await http.delete(`/challenge/question/${questionId}`);
    await fetchChallenge();
};

const listenForUpdates = async () => {
    ws = new Sockette("ws://" + window.location.host + "/ws", {
        onmessage: async (event) => {
            let data: Challenge = JSON.parse(event.data);
            if (data.id === challenge.value?.id) {
                await fetchChallenge();
            }
        },
    });
};

const updateChallenge = async () => {
    await http.patch(`/challenge/${challenge.value?.id}`, {
        name: challenge.value?.name,
        description: challenge.value?.description,
    });
};

const newQuestionsFromCaption = async () => {
    if (!challenge.value) {
        return;
    }

    questionGenerationCorpus.value = challenge.value.caption;

    await http.post(`/challenge/new-questions/${challenge.value?.id}`, {
        corpus: challenge.value.caption,
    });
};

const newQuestionsFromCorpus = async () => {
    if (questionGenerationCorpus.value === "") {
        return;
    }

    await http.post(`/challenge/new-questions/${challenge.value?.id}`, {
        corpus: questionGenerationCorpus.value,
    });
};

onBeforeMount(async () => {
    await fetchChallenge();
    await listenForUpdates();
});

onBeforeUnmount(() => {
    if (ws) {
        ws.close();
    }
});

</script>
<template>
    <form>
        <div v-if="challenge">
            <div class="pb-4">
                <h1 class="text-3xl pb-1">
                    <span class="text-gray-500 font-bold">
                    Challenge #{{ challenge.id }}:
                    </span>
                    {{ challenge.name }}</h1>
                <span class="bg-gray-100 text-gray-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded">
                    <Tippy>
                        Created {{ fuzzyDate(challenge.created_at) }}
                        <template #content>
                            {{ challenge.created_at }}
                        </template>
                    </Tippy>
                </span>
                <span class="bg-gray-100 text-gray-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded">
                    <Tippy>
                        Updated {{ fuzzyDate(challenge.updated_at) }}
                        <template #content>
                            {{ challenge.updated_at }}
                        </template>
                    </Tippy>
                </span>
            </div>
            <div>
                <div class="grid gap-4 mb-4 sm:grid-cols-3">
                    <div class="pb-4 col-span-2 flex flex-col gap-3">
                        <div>
                            <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Name</label>
                            <div>
                                <input v-model="challenge.name" type="text" name="name" id="name" class="modal-input" required />
                            </div>
                        </div>
                        <div>
                            <label for="description" class="block mb-2 text-sm font-medium text-gray-900">Description</label>
                            <div>
                                <input v-model="challenge.description" type="text" name="description" id="description" class="modal-input" required />
                            </div>
                        </div>
                        <div>
                            <label for="description" class="block mb-2 text-sm font-medium text-gray-900">Generate New Questions</label>
                            <textarea v-model="questionGenerationCorpus" placeholder="Generate new questions based on your own corpus here." class="modal-input" cols="30" rows="10"></textarea>
                        </div>
                        <div class="self-start flex gap-2">
                            <button
                                type="button"
                                @click="updateChallenge"
                                class="bg-gray-600 hover:bg-gray-700 text-white font-bold py-1 px-3 rounded flex items-center uppercase self-start"
                            >
                                Update
                            </button>
                            <button
                                @click="newQuestionsFromCorpus"
                                type="button"
                                class="bg-amber-500 hover:bg-amber-600 text-white font-bold py-1 px-3 rounded flex items-center uppercase self-start"
                            >
                                Generate
                            </button>
                            <button
                                @click="newQuestionsFromCaption"
                                type="button"
                                class="bg-amber-500 hover:bg-amber-600 text-white font-bold py-1 px-3 rounded flex items-center uppercase self-start"
                            >
                                Regenerate from caption
                            </button>
                        </div>
                    </div>
                    <div class="pb-4 col-span-1 m-6">
                        <figure>
                            <img class="w-full h-auto rounded-lg" :src="challenge.filename" alt="challenge image" />
                            <figcaption class="mt-2 text-sm text-center text-gray-500">
                                <template v-if="challenge.caption">
                                    Caption: {{ challenge.caption }}
                                </template>
                                <template v-else>
                                    <span class="italic">
                                    Caption Processing
                                    </span>
                                </template>
                            </figcaption>
                        </figure>
                    </div>
                </div>

                <h1 class="text-xl uppercase tracking-wide text-gray-500 font-bold pb-2">Questions and Answers</h1>
                <div class="space-y-5">
                    <div v-if="challenge.questions.length > 0" v-for="question in challenge.questions" class="text-base">
                        <div class="flex gap-4 items-center">
                            <div class="w-6/12">
                                <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Question</label>
                                <input v-model="question.question" type="text" name="name" id="name" class="modal-input" required />
                            </div>
                            <div class="w-6/12">
                                <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Answer</label>
                                <input v-model="question.answer" type="text" name="name" id="name" class="modal-input" required />
                            </div>
                            <div class="w-10 inline-flex justify-center items-center mt-6 gap-1">
                                <Tippy>
                                    <PencilSquareIcon @click="updateQuestion(question.id, question.question, question.answer)" class="h-6 w-6 text-gray-400 hover:text-gray-500 hover:scale-110 cursor-pointer" />
                                    <template #content>
                                        Update
                                    </template>
                                </Tippy>
                                <Tippy>
                                    <TrashIcon @click="deleteQuestion(question.id)" class="h-6 w-6 text-gray-400 hover:text-gray-500 hover:scale-110 cursor-pointer" />
                                    <template #content>
                                        Delete
                                    </template>
                                </Tippy>
                            </div>
                        </div>
                    </div>
                    <div v-else>
                        <p class="text-gray-400 text-2xl text-center bg-gray-100 py-8">No questions yet, generate some?</p>
                    </div>
                </div>

            </div>
        </div>
    </form>
</template>

<style scoped>
.modal-input {
    @apply block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-gray-600 focus:ring-gray-600;
}
</style>
