const { testConnection } = require('../config/db');

// Health check endpoint
exports.checkHealth = async (req, res, next) => {
  try {
    // Check database connection
    const dbStatus = await testConnection();
    
    res.status(200).json({
      success: true,
      status: 'OK',
      timestamp: new Date(),
      checks: {
        database: dbStatus ? 'connected' : 'disconnected',
        api: 'running'
      }
    });
  } catch (err) {
    console.error('Error in health check:', err.message);
    next(err);
  }
};