const User = require('../models/user');
const { generateToken, isValidEmail } = require('../utils/helpers');
const { MESSAGES } = require('../config/constants');
const { token } = require('morgan');

// Register a new user
exports.register = async (req, res, next) => {
  try {
    const { name, email, organization, password } = req.body;
    
    // Validate input
    if (!name || !email || !password || !organization) {
      return res.status(400).json({
        success: false,
        message: 'Please provide name, email, organization and password'
      });
    }
    
    // Validate email format
    if (!isValidEmail(email)) {
      return res.status(400).json({
        success: false,
        message: 'Please provide a valid email address'
      });
    }
    
    // Check if password meets minimum requirements
    if (password.length < 6) {
      return res.status(400).json({
        success: false,
        message: 'Password must be at least 6 characters long'
      });
    }

    if (organization.length==0) {
        return res.status(400).json({
          success: false,
          message: 'Please provide an organization name'
        });
      }
    
    // Check if user already exists
    const existingUser = await User.findByEmail(email);
    if (existingUser) {
      return res.status(400).json({
        success: false,
        message: MESSAGES.USER_EXISTS
      });
    }
    
    // Create new user
    const user = await User.create(name, email,organization, password);

    var token = generateToken(user.id)
    
    // Send response
    res.status(201).json({
      success: true,
      message: MESSAGES.USER_CREATED,
      user: {
        id: user.id,
        name: user.name,
        token: token,
        email: user.email,
        organization: user.organization,
        created_at: user.created_at
      }
    });
  } catch (err) {
    console.error('Error in register:', err.message);
    next(err);
  }
};

// Login user
exports.login = async (req, res, next) => {
  try {
    const { email, password } = req.body;
    
    // Validate input
    if (!email || !password) {
      return res.status(400).json({
        success: false,
        message: 'Please provide email and password'
      });
    }
    
    // Find user by email
    const user = await User.findByEmail(email);
    if (!user) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.INVALID_CREDENTIALS
      });
    }
    
    // Validate password
    const isPasswordValid = await User.validatePassword(password, user.password);
    if (!isPasswordValid) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.INVALID_CREDENTIALS
      });
    }

    var token = generateToken(user.id)
    
    // Send response
    res.status(200).json({
      success: true,
      message: MESSAGES.LOGIN_SUCCESS,
      user: {
        id: user.id,
        name: user.name,
        token: token,
        email: user.email,
        organization: user.organization,
        created_at: user.created_at
      }
    });
  } catch (err) {
    console.error('Error in login:', err.message);
    next(err);
  }
};

// Get current user profile
exports.getMe = async (req, res, next) => {
  try {
    const user = await User.findById(req.user.id);
    
    if (!user) {
      return res.status(404).json({
        success: false,
        message: 'User not found'
      });
    }
    
    res.status(200).json({
      success: true,
      user
    });
  } catch (err) {
    console.error('Error in getMe:', err.message);
    next(err);
  }
};
