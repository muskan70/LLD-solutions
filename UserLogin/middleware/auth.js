const { verifyToken } = require('../utils/helpers');
const { MESSAGES } = require('../config/constants');
const User = require('../models/user');

// Authentication middleware
exports.authenticate = async (req, res, next) => {
  try {
    // Get token from header
    const authHeader = req.header('Authorization');
    
    if (!authHeader) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.MISSING_AUTH_HEADER
      });
    }
    
    // Check if the header starts with 'Bearer '
    if (!authHeader.startsWith('Bearer ')) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.INVALID_TOKEN
      });
    }
    
    // Extract the token
    const token = authHeader.substring(7);
    
    // Verify token
    const decoded = verifyToken(token);
    if (!decoded) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.INVALID_TOKEN
      });
    }
    
    // Check if user exists
    const user = await User.findById(decoded.id);
    if (!user) {
      return res.status(401).json({
        success: false,
        message: MESSAGES.AUTHORIZATION_DENIED
      });
    }
    
    // Set user in request
    req.user = { id: user.id };
    next();
  } catch (err) {
    console.error('Error in authentication middleware:', err.message);
    res.status(401).json({
      success: false,
      message: MESSAGES.AUTHORIZATION_DENIED
    });
  }
};
