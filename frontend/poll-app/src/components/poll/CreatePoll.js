import { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Typography,
  Paper,
  IconButton,
  List,
  ListItem,
  ListItemSecondaryAction,
  Container,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import { useNavigate, useOutletContext } from "react-router-dom";

function CreatePoll() {
  const [question, setQuestion] = useState("");
  const [pollOptions, setPollOptions] = useState([""]); // Start with one empty option
  const [error, setError] = useState("");

  const { jwtToken } = useOutletContext();
  const { user } = useOutletContext();

  const navigate = useNavigate();

  const handleAddOption = () => {
    setPollOptions([...pollOptions, ""]);
  };

  const handleOptionChange = (index, value) => {
    const newOptions = [...pollOptions];
    newOptions[index] = value;
    setPollOptions(newOptions);
  };

  const handleRemoveOption = (index) => {
    if (pollOptions.length > 1) {
      // Ensure at least one option remains
      const newOptions = pollOptions.filter((_, i) => i !== index);
      setPollOptions(newOptions);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    // Validation
    if (!question.trim()) {
      setError("Question is required");
      return;
    }

    // Filter out empty options and check if we have at least two valid options
    const validOptions = pollOptions.filter((option) => option.trim());
    if (validOptions.length < 2) {
      setError("At least two valid options are required");
      return;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${jwtToken}`);

    const body = JSON.stringify({
      question: question.trim(),
      options: validOptions.map((option) => option.trim()),
      user_id: user.id,
    });

    console.log("BODY");
    console.log(body);

    fetch("/polls", {
      method: "POST",
      headers: headers,
      credentials: "include",
      body: body,
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.error) {
          console.error(data.message);
        } else {
          // Reset form on success
          setQuestion("");
          setPollOptions([""]);
          navigate("/polls");
        }
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <Container maxWidth="md" sx={{ p: 2 }}>
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          mx: "auto",
          mt: 4,
          p: 3,
        }}
      >
        <Paper elevation={3} sx={{ p: 3 }}>
          <Typography variant="h5" component="h1" gutterBottom>
            Create New Poll
          </Typography>

          {error && (
            <Typography color="error" sx={{ mb: 2 }}>
              {error}
            </Typography>
          )}

          <TextField
            fullWidth
            label="Question"
            value={question}
            placeholder=""
            onChange={(e) => setQuestion(e.target.value)}
            margin="normal"
            variant="filled"
            required
          />

          <Typography variant="h6" sx={{ mt: 3, mb: 2 }}>
            Poll Options
          </Typography>

          <List>
            {pollOptions.map((option, index) => (
              <ListItem key={index} disablePadding sx={{ mb: 1 }}>
                <TextField
                  fullWidth
                  label={`Option ${index + 1}`}
                  value={option}
                  onChange={(e) => handleOptionChange(index, e.target.value)}
                  variant="filled"
                  size="small"
                  required
                />
                <ListItemSecondaryAction>
                  <IconButton
                    edge="end"
                    onClick={() => handleRemoveOption(index)}
                    disabled={pollOptions.length <= 1}
                  >
                    <DeleteIcon />
                  </IconButton>
                </ListItemSecondaryAction>
              </ListItem>
            ))}
          </List>

          <Button
            startIcon={<AddCircleIcon />}
            onClick={handleAddOption}
            sx={{ mt: 2 }}
          >
            Add Option
          </Button>

          <Box sx={{ mt: 4 }}>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              fullWidth
              size="large"
            >
              Create Poll
            </Button>
          </Box>
        </Paper>
      </Box>
    </Container>
  );
}

export default CreatePoll;
