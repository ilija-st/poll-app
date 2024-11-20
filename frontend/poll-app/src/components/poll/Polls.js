import { Box } from "@mui/material";
import Poll from "./Poll";
import { useEffect, useState } from "react";

const Polls = () => {
  const [polls, setPolls] = useState([]);

  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(`/polls`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setPolls(data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  // Parse polls data
  if (polls) {
    polls.map((poll) => {
      poll.poll_options = Object.values(poll.edges.poll_options);
      poll.author =
        poll.edges.user.first_name + " " + poll.edges.user.last_name;
      poll.total_votes = 0;
      if (poll.poll_options) {
        poll.options_count = poll.poll_options.length;
        poll.poll_options.map((pollOpt) => {
          pollOpt.votes = pollOpt.edges.votes;
          pollOpt.num_votes = pollOpt.votes?.length;
          if (pollOpt.votes?.length) {
            poll.total_votes += pollOpt.votes.length;
          }
        });
      }
      console.log(poll);
    });
  }

  const handleVote = async (pollId, optionId) => {
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 1000));
    console.log("Voted on poll:", pollId, "with option:", optionId);
  };

  const handleEdit = (pollId) => {
    console.log("Editing poll:", pollId);
    // Navigate to edit page or open edit modal
  };

  const handleDelete = (pollId) => {
    console.log("Deleting poll:", pollId);
    // Show confirmation dialog and delete poll
  };

  return (
    <Box sx={{ p: 2 }}>
      {polls.map((poll) => (
        <Poll
          key={poll.id}
          poll={poll}
          onVote={(id) => console.log("Vote:", id)}
          onEdit={(id) => console.log("Edit:", id)}
          onDelete={(id) => console.log("Delete:", id)}
          isOwner={true}
          hasVoted={false}
        />
      ))}
    </Box>
  );
};

export default Polls;
