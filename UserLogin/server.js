const express = require('express');
const helmet = require('helmet');
const morgan = require('morgan');
const cors = require('cors');
const { initializeDatabase } = require('./config/db');
const { register, login, getMe} = require('./controllers/authController');
const { checkHealth } = require('./controllers/healthController');
const { authenticate } = require('./middleware/auth');

// Initialize express app
const app = express();

// Middleware
app.use(helmet()); // Security headers
app.use(morgan('dev')); // Logging
app.use(cors()); // Enable CORS
app.use(express.json()); // Parse JSON bodies
app.use(express.urlencoded({ extended: true })); // Parse URL-encoded bodies

// Initialize database
initializeDatabase();

app.post('/api/user/register', register);
app.post('/api/user/login', login);
app.get('/status', checkHealth);
app.get('/dashboard', authenticate, getMe);

// Error handling middleware
app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(err.status || 500).json({
    success: false,
    message: err.message || 'Internal Server Error',
    error: process.env.NODE_ENV === 'production' ? {} : err
  });
});

// Start server
const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});

module.exports = app;