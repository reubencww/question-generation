<template>
    <div class="upload">
        <div class="grid grid-cols-2 pb-4 gap-4">
            <div>
                <label
                    for="name"
                    class="block mb-2 text-sm font-medium text-gray-900"
                    >Name</label
                >
                <div>
                    <input
                        v-model="form.name"
                        type="text"
                        name="name"
                        id="name"
                        class="modal-input"
                        required
                    />
                </div>
            </div>
            <div>
                <label
                    for="description"
                    class="block mb-2 text-sm font-medium text-gray-900"
                    >Description</label
                >
                <div>
                    <input
                        v-model="form.description"
                        type="text"
                        name="description"
                        id="description"
                        class="modal-input"
                        required
                    />
                </div>
            </div>
        </div>
        <FilePondComponent
            ref="filePondRef"
            label-idle="Drop file here..."
            accepted-file-types="image/jpeg, image/png"
            :allow-multiple="false"
            credits=""
        />
        <button
            class="bg-gray-600 hover:bg-gray-700 text-white text-sm font-bold py-2 px-3 rounded flex items-center uppercase self-start"
            @click="upload"
        >
            Upload
        </button>
    </div>
</template>

<script lang="ts" setup>
import axios from "axios";
import type { FilePond } from "filepond";
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";
import FilePondPluginImagePreview from "filepond-plugin-image-preview";
import "filepond-plugin-image-preview/dist/filepond-plugin-image-preview.min.css";
import "filepond/dist/filepond.min.css";
import { ref } from "vue";
import vueFilePond from "vue-filepond";
import { useRouter } from "vue-router";

import { Challenge } from "@/challenges";

interface ChallengeResponse {
    data: Challenge;
}

const router = useRouter();

// Create component
const FilePondComponent = vueFilePond(
    FilePondPluginFileValidateType,
    FilePondPluginImagePreview
);

const filePondRef = ref<FilePond | null>(null);
const form = ref({
    name: "",
    description: "",
});

const upload = async () => {
    if (!filePondRef.value) {
        return;
    }

    let file = filePondRef.value.getFile();

    let formData = new FormData();
    formData.append("name", form.value.name);
    formData.append("description", form.value.description);
    formData.append("image", file.file);

    let response = await axios.post<ChallengeResponse>(
        "/api/v1/challenge",
        formData
    );
    await router.push({
        name: "challenge.show",
        params: { id: response.data.data.id },
    });
};
</script>

<style scoped>
.modal-input {
    @apply block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-gray-600 focus:ring-gray-600;
}
</style>
