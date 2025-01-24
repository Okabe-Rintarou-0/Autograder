export async function urlToFile(url: string, fileName: string) {
    const response = await fetch(url);
    const blob = await response.blob();
    const file = new File([blob], fileName, { type: blob.type });

    return file;
}