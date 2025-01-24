export interface User {
    id: number;
    name: string;
    created_at: string;
    sortable_name: string;
    short_name: string;
    login_id: string;
    email: string | null;
}