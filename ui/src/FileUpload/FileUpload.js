import http from '../http-common'
import {useState} from "react";

// https://bezkoder.com/react-hooks-file-upload/

export const UploadUI = () => {
    // define states using react hooks
    const [selectedFiles, setSelectedFiles] = useState(undefined);
    const [currentFile, setCurrentFile] = useState(undefined);
    const [message, setMessage] = useState("");
    const [fileInfo, setFileInfo] = useState([]);

    const selectFile = (event) => {
        setSelectedFiles(event.target.file);
    };

    setCurrentFile(selectedFiles);
    upload(currentFile)
        .then((response) => {
            setMessage(response.data.message);
        })
        .then((files) => {
            setFileInfo(files.data);
        })
        .catch(() => {
            setMessage('could not upload file');
            setCurrentFile(undefined);
        });
    setSelectedFiles(undefined);

    return (
        <div>
            <label className="btn btn-default">
                <input type="file" onChange={selectFile} />
            </label>

            <button
                className="btn btn-success"
                disabled={!selectedFiles}
                onClick={upload}
            >Upload</button>

            <div className='alert alert-dark' role='alert'>
                {message}
            </div>
        </div>
    );
};

const upload = (file) => {
    const formData = new FormData();
    formData.append('file', file);

    return http.post('/upload', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        },
    }).then((resp) => {
        if(resp.status === 200) {
            console.log('file uploaded')
        }
    })
};

export default {
    UploadUI,
};