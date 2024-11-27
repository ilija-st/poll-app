import { useState, useEffect } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  IconButton,
  List,
  ListItem,
  ListItemSecondaryAction,
  Typography,
  Alert,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import { useNavigate, useOutletContext } from "react-router-dom";

function EditPollDialog({ open, handleClose, poll }) {
  const [question, setQuestion] = useState("");
  const [pollOptions, setPollOptions] = useState([]);
  const [newOptions, setNewOptions] = useState([""]); // For new options
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const { jwtToken } = useOutletContext();

  const navigate = useNavigate();

  // Initialize form with poll data when opened
  useEffect(() => {
    if (poll) {
      setQuestion(poll.question);
      setPollOptions(poll.poll_options);
      setNewOptions([""]); // Reset new options
      setError("");
    }
  }, [poll]);

  const handleNewOptionChange = (index, value) => {
    const updatedNewOptions = [...newOptions];
    updatedNewOptions[index] = value;
    setNewOptions(updatedNewOptions);
  };

  const handleAddNewOption = () => {
    setNewOptions([...newOptions, ""]);
  };

  const handleRemoveNewOption = (index) => {
    const updatedNewOptions = [...newOptions];
    updatedNewOptions.splice(index, 1);
    setNewOptions(updatedNewOptions);
  };

  const handleSubmit = () => {
    setError("");
    setLoading(true);

    const validExistingOptions = pollOptions;
    const validNewOptions = newOptions.filter((option) => option.trim());

    if (validExistingOptions.length + validNewOptions.length < 2) {
      setError("At least two options are required");
      setLoading(false);
      return;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${jwtToken}`);

    const updateData = JSON.stringify({
      poll_id: poll.id,
      options: [...validNewOptions.map((option) => option.trim())],
    });

    fetch(`${process.env.REACT_APP_BACKEND}/polls`, {
      method: "PUT",
      headers: headers,
      credentials: "include",
      body: updateData,
    })
      .then((res) => {
        if (res.status !== 200) {
          throw new Error("Failed to update poll");
        }
        return res.json();
      })
      .then((data) => {
        handleClose(true); // Pass true to indicate successful update
        navigate("/polls");
      })
      .catch((err) => {
        setError(err.message);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <Dialog
      open={open}
      onClose={() => handleClose(false)}
      maxWidth="sm"
      fullWidth
    >
      <DialogTitle>Edit Poll</DialogTitle>
      <DialogContent>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        <Typography variant="h5" sx={{ mt: 2, mb: 1 }}>
          {question}
        </Typography>

        <Typography variant="h6" sx={{ mt: 3, mb: 1 }}>
          Existing Options
        </Typography>

        <List>
          {pollOptions.map((option) => (
            <ListItem key={option.id} sx={{ px: 0 }}>
              <TextField
                fullWidth
                value={option.title}
                disabled={true}
                variant="outlined"
                size="small"
              />
            </ListItem>
          ))}
        </List>

        <Typography variant="h6" sx={{ mt: 3, mb: 1 }}>
          New Options
        </Typography>

        <List>
          {newOptions.map((option, index) => (
            <ListItem key={`new-${index}`} sx={{ px: 0 }}>
              <TextField
                fullWidth
                label={`New Option ${index + 1}`}
                value={option}
                onChange={(e) => handleNewOptionChange(index, e.target.value)}
                variant="filled"
                size="small"
              />
              <ListItemSecondaryAction>
                <IconButton
                  edge="end"
                  onClick={() => handleRemoveNewOption(index)}
                >
                  <DeleteIcon />
                </IconButton>
              </ListItemSecondaryAction>
            </ListItem>
          ))}
        </List>

        <Button
          startIcon={<AddCircleIcon />}
          onClick={handleAddNewOption}
          sx={{ mt: 1 }}
        >
          Add New Option
        </Button>
      </DialogContent>

      <DialogActions sx={{ px: 3, pb: 2 }}>
        <Button onClick={() => handleClose(false)}>Cancel</Button>
        <Button onClick={handleSubmit} variant="contained" disabled={loading}>
          {loading ? "Saving..." : "Save Changes"}
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default EditPollDialog;
