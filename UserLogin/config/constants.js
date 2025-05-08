module.exports = {
	// Database
	DB_USER_TABLE: 'users',
	
	// JWT
	JWT_SECRET: process.env.JWT_SECRET || 'your_jwt_secret_key',
	JWT_EXPIRES_IN: '24h',
	
	// Password hashing
	SALT_ROUNDS: 10,
	
	// Response messages
	MESSAGES: {
	  USER_CREATED: 'User registered successfully',
	  USER_EXISTS: 'User already exists with this email',
	  LOGIN_SUCCESS: 'Login successful',
	  INVALID_CREDENTIALS: 'Invalid email or password',
	  SERVER_ERROR: 'Server error',
	  AUTHORIZATION_DENIED: 'Authorization denied',
	  INVALID_TOKEN: 'Invalid token',
	  MISSING_AUTH_HEADER: 'No authorization token, access denied'
	}
  };