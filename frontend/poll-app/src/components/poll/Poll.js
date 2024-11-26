import React, { useEffect, useState } from "react";
import {
  Card,
  CardContent,
  CardActions,
  Typography,
  Button,
  Chip,
  Box,
  IconButton,
  Tooltip,
  Container,
  Collapse,
  Divider,
  RadioGroup,
  FormControlLabel,
  Radio,
  CircularProgress,
} from "@mui/material";
import {
  HowToVote as VoteIcon,
  Person as PersonIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  ExpandMore as ExpandMoreIcon,
  Delete,
  Close,
  Stop,
} from "@mui/icons-material";
import { formatDistance } from "date-fns";
import { useOutletContext } from "react-router-dom";
import VotersModal from "../voters-modal/VotersModal";
import EditPollDialog from "./EditPoll";

const Poll = ({ poll, onVote, onClosePoll, onDelete, isOwner }) => {
  const [isExpanded, setIsExpanded] = useState(false);
  const [selectedOption, setSelectedOption] = useState("");
  const [opt, setOpt] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);
  const [editDialogOpen, setEditDialogOpen] = useState(false);

  const {
    id,
    question,
    created_at,
    voted,
    total_votes = 0,
    author = "",
    options_count = 0,
    poll_options = [],
  } = poll;

  const pollAuthor = isOwner ? "You" : author;

  // Format the creation date relative to now (e.g., "2 hours ago")
  const timeAgo = formatDistance(new Date(created_at), new Date(), {
    addSuffix: true,
  });

  const handleVoteClick = () => {
    if (!isExpanded) {
      if (voted) {
        setSelectedOption(poll.voted_option_id);
      }
    }
    setIsExpanded(!isExpanded);
  };

  const handleVotersClick = (option) => {
    setOpt(option);
    setModalOpen(true);
  };

  const handleOptionChange = (event) => {
    setSelectedOption(event.target.value);
  };

  const handleVoteSubmit = async () => {
    if (!selectedOption) return;

    setIsSubmitting(true);
    try {
      onVote(id, selectedOption);
      setIsExpanded(false);
    } catch (error) {
      console.error("Voting failed:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleEditClose = (wasUpdated) => {
    setEditDialogOpen(false);
    if (wasUpdated) {
      // TODO: Refresh poll data
    }
  };

  return (
    <Container maxWidth="lg" sx={{ p: 2 }}>
      <Card
        sx={{
          minWidth: 275,
          transition: "transform 0.2s",
          "&:hover": {
            transform: isExpanded ? "none" : "translateY(-2px)",
            boxShadow: 4,
          },
        }}
      >
        <CardContent>
          {/* Author and time */}
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              mb: 2,
              color: "text.secondary",
            }}
          >
            <PersonIcon sx={{ fontSize: 20, mr: 1 }} />
            <Typography variant="body2" component="span">
              {pollAuthor} • {timeAgo}
            </Typography>
          </Box>

          {/* Poll question */}
          <Typography
            variant="h6"
            component="div"
            sx={{
              mb: 2,
              overflow: "hidden",
              textOverflow: "ellipsis",
              display: "-webkit-box",
              WebkitLineClamp: 2,
              WebkitBoxOrient: "vertical",
            }}
          >
            {question}
          </Typography>

          {/* Stats */}
          <Box sx={{ display: "flex", gap: 1, mb: 3 }}>
            <Chip
              icon={<VoteIcon />}
              label={`${total_votes} votes`}
              size="small"
              color="primary"
              variant="outlined"
            />
            <Chip
              label={`${options_count} options`}
              size="small"
              variant="outlined"
            />
          </Box>
        </CardContent>

        {/* Expandable Voting Options */}
        <Collapse in={isExpanded} timeout="auto" unmountOnExit>
          <Divider />
          <CardContent sx={{ pt: 2, pb: 1 }}>
            <RadioGroup value={selectedOption} onChange={handleOptionChange}>
              {poll_options.map((option) => (
                <FormControlLabel
                  key={option.id}
                  value={option.id.toString()}
                  control={<Radio disabled={voted} />}
                  label={
                    <Box
                      sx={{
                        display: "flex",
                        gap: 2,
                        justifyContent: "space-between",
                        width: "100%",
                        alignItems: "center",
                      }}
                    >
                      <Typography>{option.title}</Typography>
                      {voted && (
                        <Button
                          onClick={() => handleVotersClick(option)}
                          variant="text"
                          size="small"
                        >
                          {option.num_votes || "No "} votes
                        </Button>
                      )}
                    </Box>
                  }
                  sx={{
                    margin: "8px 0",
                    padding: "8px",
                    width: "100%",
                    borderRadius: 1,
                    "&:hover": {
                      backgroundColor: "action.hover",
                    },
                  }}
                />
              ))}
            </RadioGroup>
          </CardContent>
          <CardActions sx={{ justifyContent: "flex-end", p: 2 }}>
            <Button onClick={() => setIsExpanded(false)} sx={{ mr: 1 }}>
              Cancel
            </Button>
            <Button
              variant="contained"
              onClick={handleVoteSubmit}
              disabled={voted || !selectedOption}
            >
              {isSubmitting ? <CircularProgress size={24} /> : "Submit Vote"}
            </Button>
          </CardActions>
        </Collapse>
        <VotersModal
          open={modalOpen}
          handleClose={() => setModalOpen(false)}
          option={opt}
          optionTitle={opt.title}
        />

        {/* Main Actions */}
        <CardActions
          sx={{
            display: "flex",
            justifyContent: "space-between",
            px: { xs: 2, sm: 3 },
            pb: { xs: 2, sm: 3 },
          }}
        >
          <Box
            sx={{
              display: "flex",
              gap: 2,
              flexGrow: 1,
              maxWidth: isOwner ? "calc(100% - 100px)" : "100%",
            }}
          >
            <Button
              variant="contained"
              startIcon={<VoteIcon />}
              onClick={handleVoteClick}
              color="info"
              size="large"
              endIcon={
                <ExpandMoreIcon
                  sx={{
                    transform: isExpanded ? "rotate(180deg)" : "rotate(0deg)",
                    transition: "transform 0.3s",
                  }}
                />
              }
              sx={{
                flexGrow: 1,
                py: 1,
              }}
            >
              {voted ? "Voted" : "Vote Now"}
            </Button>
          </Box>

          {/* Edit and Delete buttons */}
          {isOwner && (
            <Box sx={{ display: "flex", gap: 1 }}>
              <Tooltip title="Close Poll">
                <IconButton size="medium" onClick={() => onClosePoll(poll)}>
                  <Stop />
                </IconButton>
              </Tooltip>
              <Tooltip title="Edit Poll">
                <IconButton
                  size="medium"
                  onClick={() => setEditDialogOpen(true)}
                >
                  <EditIcon />
                </IconButton>
              </Tooltip>
              <EditPollDialog
                open={editDialogOpen}
                handleClose={handleEditClose}
                poll={poll}
              />
              <Tooltip title="Delete Poll">
                <IconButton
                  size="medium"
                  onClick={() => onDelete(poll)}
                  color="error"
                >
                  <DeleteIcon />
                </IconButton>
              </Tooltip>
            </Box>
          )}
        </CardActions>
      </Card>
    </Container>
  );
};

export default Poll;
