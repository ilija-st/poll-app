import {
  Modal,
  Box,
  Typography,
  List,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Avatar,
  IconButton,
} from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";
import PersonIcon from "@mui/icons-material/Person";
import { useState, useEffect } from "react";

function VotersModal({ open, handleClose, option }) {
  const [voters, setVoters] = useState([]);

  // Fetch voters when modal opens
  useEffect(() => {
    if (open && option.num_votes && option.num_votes > 0) {
      setVoters(option.edges.votes);
    }
  }, [open]);

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="voters-modal-title"
    >
      <Box
        sx={{
          position: "absolute",
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -50%)",
          width: { xs: "90%", sm: 400 },
          bgcolor: "background.paper",
          borderRadius: 2,
          boxShadow: 24,
          p: 4,
          maxHeight: "80vh",
          overflow: "auto",
        }}
      >
        {/* Header */}
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            mb: 2,
          }}
        >
          <Typography id="voters-modal-title" variant="h6" component="h2">
            Voters for "{option.title}"
          </Typography>
          <IconButton onClick={handleClose} size="small" sx={{ ml: 2 }}>
            <CloseIcon />
          </IconButton>
        </Box>

        {/* Content */}
        {option.num_votes === 0 ? (
          <Typography
            sx={{ py: 2, textAlign: "center", color: "text.secondary" }}
          >
            No votes yet for this option
          </Typography>
        ) : (
          <List sx={{ pt: 0 }}>
            {voters.map((voter) => (
              <ListItem key={voter.id}>
                <ListItemAvatar>
                  <Avatar>
                    <PersonIcon />
                  </Avatar>
                </ListItemAvatar>
                <ListItemText
                  primary={
                    voter.edges.user.first_name +
                    " " +
                    voter.edges.user.last_name
                  }
                  secondary={`Voted on ${new Date(
                    voter.created_at
                  ).toLocaleDateString()}`}
                />
              </ListItem>
            ))}
          </List>
        )}
      </Box>
    </Modal>
  );
}

export default VotersModal;
