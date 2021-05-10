import http from "../http-common";

const upload = (file, onUploadProgress) => {
    const formData = new FormData();
    formData.append('file', file);

    return http.post('/upload', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        onUploadProgress,
    }).then((resp) => {
        if(resp.status === 200) {
            console.log('file uploaded')
        }
    })
}

const getFile = () => {
    return http.get("/file");
}

export default {
    upload,
    getFile
}