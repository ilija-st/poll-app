import { Box } from "@mui/material";
import Poll from "./Poll";

const Polls = () => {
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
  const polls = [
    {
      id: 1,
      question: "What's your favorite programming language?",
      created_at: "2024-03-19T10:00:00Z",
      total_votes: 42,
      author: "John Doe",
      options_count: 5,
    },
    {
      id: 2,
      question: "What's your most hated programming language?",
      created_at: "2023-03-19T10:00:00Z",
      total_votes: 1232,
      author: "Jon Snow",
      options_count: 6,
    },
    // ... more polls
  ];

  return (
    <Box sx={{ p: 2 }}>
      {polls.map((poll) => (
        <Poll
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
