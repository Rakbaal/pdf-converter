export default function convertFileList(list: FileList): File[] {
    const files: File[] = []
    for (let i = 0; i < list.length; i++) {
        files.push(list[i])
    }

    return files
}