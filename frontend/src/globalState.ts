import { writable } from "svelte/store";

type CurrentPage = { page: 'home' } | { page: 'list-files' } | { page: 'view-file', fileName: string }

export const currentPage = writable<CurrentPage>({ page: 'home' });

export function goToPage(page: CurrentPage) {
  currentPage.set(page);
}
