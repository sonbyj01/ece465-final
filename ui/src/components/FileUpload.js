import React, { useState, useEffect } from "react";
import UploadService from "../services/FileUploadService";

const UploadFile = () => {
    // declare and initialize react hooks
    const [selectedFile, setSelectedFile] = useState(undefined);
    const [progress, setProgress] = useState(0);
    const [message, setMessage] = useState("");
    const [fileInfo, setFileInfo] = useState([]);

    // function selectFile to get selected files from input element
    const selectFile = (event) => {
        setSelectedFile(event.target.files[0])
    };

    const upload = () => {
        setProgress(0);
        UploadService.upload(selectedFile, (event) => {
            setProgress(Math.round((100 * event.loaded) / event.total));
        })
            .then((response) => {
                setMessage(response.data.message);
                return UploadService.getFile();
            })
            .then((file) => {
                setFileInfo(file.data);
            })
            .catch(() => {
                setProgress(0);
                setMessage("couldn't upload file");
            });
        setSelectedFile(undefined);
    }

    useEffect(() => {
        UploadService.getFile().then((response) => {
            setFileInfo(response.data);
        });
    }, []);

    return (
        <div>
            {selectedFile && (
                <div className={"progress"}>
                    <div
                        className={"progress-bar progress-bar-info progress-bar-striped"}
                        role={"progressbar"}
                        aria-valuenow={progress}
                        aria-valuemin={"0"}
                        aria-valuemax={"100"}
                        style={{ width: progress + "%" }}
                        >
                        {progress}%
                    </div>
                </div>
            )}

            <label className={"btn btn-default"}>
                <input type={"file"} onChange={selectFile} />
            </label>

            <button
                className={"btn btn-success"}
                disabled={!selectedFile}
                onClick={upload}
                >
                Upload
            </button>

            <div className={"alert alert-dark"} role={"alert"}>
                {message}
            </div>

            <div className="card">
                <div className="card-header">List of Files</div>
                <ul className="list-group list-group-flush">
                    {fileInfo &&
                    fileInfo.map((file, index) => (
                        <li className="list-group-item" key={index}>
                            <a href={file.url}>{file.name}</a>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default UploadFile;