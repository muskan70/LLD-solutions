const { pool } = require('../config/db');
const { DB_USER_TABLE } = require('../config/constants');
const bcrypt = require('bcryptjs');

class User {
  // Create a new user
  static async create(name, email, organization, password) {
    const hashedPassword = await bcrypt.hash(password, 10);
    
    const query = {
      text: `INSERT INTO ${DB_USER_TABLE} (name, email, organization, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, organization, created_at`,
      values: [name, email.toLowerCase(),organization, hashedPassword]
    };
    
    const result = await pool.query(query);
    return result.rows[0];
  }
  
  // Find user by email
  static async findByEmail(email) {
    const query = {
      text: `SELECT * FROM ${DB_USER_TABLE} WHERE email = $1`,
      values: [email.toLowerCase()]
    };
    
    const result = await pool.query(query);
    return result.rows[0];
  }
  
  // Find user by ID
  static async findById(id) {
    const query = {
      text: `SELECT id, name, email, created_at FROM ${DB_USER_TABLE} WHERE id = $1`,
      values: [id]
    };
    
    const result = await pool.query(query);
    return result.rows[0];
  }
  
  // Validate password
  static async validatePassword(plainPassword, hashedPassword) {
    return await bcrypt.compare(plainPassword, hashedPassword);
  }
}

module.exports = User;
