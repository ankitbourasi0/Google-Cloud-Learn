//Reference: https://www.npmjs.com/package/@google-cloud/cloud-sql-connector

import express from 'express';
import mysql from 'mysql2/promise';
import { Connector } from '@google-cloud/cloud-sql-connector';
const app = express();

const connector = new Connector();
const PORT = process.env.PORT || 3000;


const getPool = async () => {
    console.log('Instance:', process.env.INSTANCE_CONNECTION_NAME);
    console.log('DB Name:', process.env.DB_NAME);

    const clientOpts = await connector.getOptions({
        instanceConnectionName: process.env.INSTANCE_CONNECTION_NAME,
        ipType: 'PUBLIC' //keep it private if use VPC
    })
    const pool = await mysql.createPool({
        ...clientOpts,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME,
    })
    return pool;
};

app.get('/health/db', async (req, res) => {
    try {
        const pool = await getPool();
        const connection = await pool.getConnection();

        const [result] = await connection.query(`SELECT NOW();`)
        console.table(result); //prints returned time value from server
        res.send("Welcome to Cloud SQL, Node.js Successfully!")
    } catch (error) {
        res.status(500).send('Database connection failed: ' + error.message)
    }
})

app.listen(PORT, () => console.log(`Server is running on PORT: ${PORT}`))