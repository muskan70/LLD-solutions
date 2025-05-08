const { Pool } = require('pg');
const { DB_USER_TABLE } = require('./constants');

// Create a new Pool instance
const pool = new Pool({
  user: 'postgres',
  host: 'localhost',
  database: 'postgres',
  password: 'postgres',
  port: 5432,
});

// Test database connection
const testConnection = async () => {
  try {
    const client = await pool.connect();
    console.log('Successfully connected to PostgreSQL database');
    client.release();
    return true;
  } catch (err) {
    console.error('Database connection error:', err.message);
    return false;
  }
};

// Initialize database with tables
const initializeDatabase = async () => {
  try {
    const client = await pool.connect();
    
    // Create users table if it doesn't exist
    await client.query(`
      CREATE TABLE IF NOT EXISTS ${DB_USER_TABLE} (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        organization VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      )
    `);
    
    console.log('Database initialized with required tables');
    client.release();
  } catch (err) {
    console.error('Error initializing database:', err.message);
  }
};

module.exports = {
  pool,
  testConnection,
  initializeDatabase,
  query: (text, params) => pool.query(text, params)
};