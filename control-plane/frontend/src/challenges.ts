export interface Challenge {
    id: number;
    created_at: string;
    updated_at: string;
    name: string;
    description: string;
    filename: string;
    caption: string;
    completed_at: string | null;
    questions: {
        id: number;
        created_at: string;
        updated_at: string;
        question: string;
        answer: string;
        challenge_id: number;
    }[];
}
