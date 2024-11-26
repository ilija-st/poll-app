import { Box } from "@mui/material";
import Poll from "./Poll";
import { useEffect, useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";
import Swal from "sweetalert2";

const Polls = () => {
  const [polls, setPolls] = useState([]);

  const { user } = useOutletContext();
  const { jwtToken } = useOutletContext();

  const navigate = useNavigate();

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
  // FIXME: REFACTOR THIS
  if (polls) {
    polls.map((poll) => {
      poll.poll_options = Object.values(poll.edges.poll_options);
      poll.author =
        poll.edges.user.first_name + " " + poll.edges.user.last_name;
      poll.user_id = poll.edges.user.id;
      // Total votes on a poll
      poll.total_votes = 0;
      // List of vote ids
      poll.votes = [];
      // Bool representing if the current user voted on a poll
      poll.voted = false;
      if (poll.poll_options) {
        poll.options_count = poll.poll_options.length;
        poll.poll_options.map((pollOpt) => {
          pollOpt.votes = pollOpt.edges?.votes;
          pollOpt.num_votes = pollOpt.votes?.length || 0;
          if (pollOpt.votes?.length) {
            pollOpt.votes.map((vote) => {
              if (vote.edges.user.id === user.id) {
                poll.voted = true;
                poll.voted_option_id = pollOpt.id;
              }
              poll.votes.push(vote.edges.user.id);
            });
            poll.total_votes += pollOpt.votes.length;
          }
        });
      }
      console.log(poll);
    });
  }

  const handleVote = (pollId, optionId) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${jwtToken}`);

    let requestOptions = {
      body: JSON.stringify({
        user_id: user.id,
        poll_id: pollId,
        poll_option_id: parseInt(optionId),
      }),
      method: "PUT",
      headers: headers,
      credentials: "include",
    };

    fetch(`/polls/${pollId}`, requestOptions)
      .then((response) => {
        if (response.status === 400) {
          Swal.fire({
            title: "Error",
            text: "You already voted on this poll.",
            icon: "error",
            confirmButtonText: "OK",
          });
        } else {
          Swal.fire({
            title: "Info",
            text: "You successfully voted on a poll!",
            icon: "info",
            confirmButtonText: "OK",
          });
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
        setPolls(polls);
      })
      .catch((err) => {
        console.error("ERROR: " + err);
      });
  };

  const handleClosePoll = (pollToClose) => {
    Swal.fire({
      title: "Not yet implemented... :(",
      text: "Come back later.",
      icon: "info",
      confirmButtonText: "OK",
    });
    // if (pollToDelete.user_id != user.id) {
    //   Swal.fire({
    //     title: "Error!",
    //     text: "You cant delete other users polls!",
    //     icon: "error",
    //     confirmButtonText: "OK",
    //   });
    // } else {
    //   Swal.fire({
    //     title: "Warning!",
    //     text: "Are you sure you want to close this poll?",
    //     icon: "warning",
    //     showCancelButton: true,
    //     confirmButtonText: "Yes",
    //     cancelButtonText: "No",
    //   }).then((result) => {
    //     if (result.isConfirmed) {
    //       Swal.fire("Closed!", "", "success");
    //     }
    //   });
    // }
  };

  const handleDelete = (pollToDelete) => {
    if (pollToDelete.user_id != user.id) {
      Swal.fire({
        title: "Error!",
        text: "You cant delete other users polls!",
        icon: "error",
        confirmButtonText: "OK",
      });
    } else {
      Swal.fire({
        title: "Warning!",
        text: "Are you sure you want to delete this poll?",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Yes",
        cancelButtonText: "No",
      }).then((result) => {
        if (result.isConfirmed) {
          Swal.fire("Deleted!", "", "success");
          // Delete poll
          const headers = new Headers();
          headers.append("Content-Type", "application/json");
          headers.append("Authorization", `Bearer ${jwtToken}`);

          const requestOptions = {
            method: "DELETE",
            credentials: "include",
            headers: headers,
            body: JSON.stringify({
              id: pollToDelete.id,
              user_id: user.id,
            }),
          };

          fetch(`/polls`, requestOptions).catch((error) => {
            console.log("error deleting a poll", error);
          });

          navigate("/polls");
        }
      });
    }
    // Show confirmation dialog and delete poll
  };

  return (
    <Box sx={{ p: 2 }}>
      {polls.map((poll) => (
        <Poll
          key={poll.id}
          poll={poll}
          onVote={(pollId, optionId) => handleVote(pollId, optionId)}
          onClosePoll={(poll) => handleClosePoll(poll)}
          onDelete={(poll) => handleDelete(poll)}
          isOwner={user.id === poll.user_id}
        />
      ))}
    </Box>
  );
};

export default Polls;
