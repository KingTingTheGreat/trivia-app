"use client";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import TextField from '@mui/material/TextField';

const HomeForm = () => {
    const [name, setName] = useState("");
    const [errorMessage, setErrorMessage] = useState("")
    const router = useRouter();

    const handleSubmit = async() => {
	fetch(`/api/player?name=${name}`, {
	    method: "POST", 
	}).then(res => res.text()).then(msg => {
		if (msg == "success") {
		    router.push(`/${name}`)
		}
	}).catch(err => {
		console.error(err);
	})
    }

    useEffect(() => {
	const getName = async() => {
	    const res = await fetch("/api/player")
	    if (res.status === 200) {
		setName(await res.text());
	    } else {
		// alert(await res.text());
		console.log(await res.text())
	    }
	}
	getName();
    }, [])

    return (
	<main className="flex flex-col items-center justify-center h-screen">
	    <h1 className="text-2xl font-semibold m-4">Enter your name: {name}</h1>
	    <form 
		className="w-fit flex justify-center items-center m-2"
		onSubmit={async(e) => {
		    console.log("submitting");
		    e.preventDefault();
		    await handleSubmit();
		}}
	    >
		<TextField variant="outlined" label="Name" value={name} required onChange={(e)=>{
		    setName(e.target.value);
		    setErrorMessage("");
		}} />
		<button 
		    className="p-4 m-2 text-4xl aspect-square flex justify-center items-center rounded-lg"
		    style={{ backgroundColor: name === "" ? "gray" : "lightgreen" }}
		    type="submit"
		    disabled={name === ""}
		>
		    <ArrowForwardIcon />
		</button>
	    </form>
	    <p
		style={{
		    visibility: errorMessage.length > 0 ? "visible" : "hidden",
		    color: "red",
		}}
	    >
		{errorMessage}
	    </p>
	</main>
    )
}

export default HomeForm;
