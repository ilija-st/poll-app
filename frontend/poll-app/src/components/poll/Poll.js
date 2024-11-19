import React, { useEffect, useState } from "react";

const Poll = (props) => {
  const [polls, setPolls] = useState([
    {
      id: 1,
      title: "Favorite Programming Language",
      description: "What's your go-to programming language?",
      options: ["JavaScript", "Python", "Java", "Ruby"],
      votes: 145,
      active: true,
      created: "2024-03-15",
    },
    {
      id: 2,
      title: "Best Frontend Framework",
      description: "Which frontend framework do you prefer?",
      options: ["React", "Vue", "Angular", "Svelte"],
      votes: 89,
      active: true,
      created: "2024-03-14",
    },
  ]);

  useEffect(() => {
    setPolls(polls);
    console.log("Use effect");
  }, [polls]);

  return (
    <>
      <h1>Polls</h1>
    </>
  );
};

export default Poll;
