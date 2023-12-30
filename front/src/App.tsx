import React, { ChangeEvent, useEffect, useState } from 'react';
import './App.css';
import convertFileList from './convertFileList';

function App() {
  const [files, setFile] = useState<File[] | null>(null)

  const handleFileChange = async (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const newList = convertFileList(e.target.files)
      setFile([...newList])
    }
  }
  
  const handleSubmit = async (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault()
    if (files) {
      const formData = new FormData()
      for (const index in files) {
        formData.append(`files`, files[index])
      }
      
      await fetch("http://127.0.0.1:8910/files", {
        method: "POST",
        body: formData
      })
    }
  }

  const handleTrigger = async (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault()
    await fetch("http://127.0.0.1:8910/trigger")
  }

  return (
    <div className="App">
      <input type='file' accept='.png, .jpg, .jpeg, .pdf, .webp, .gif' onChange={(e) => handleFileChange(e)} multiple/>
      <button type='submit' onClick={(e) => handleSubmit(e)}>Valider</button>
      <button type='submit' onClick={(e) => handleTrigger(e)} >Trigger</button>
    </div>
  );
}

export default App;
