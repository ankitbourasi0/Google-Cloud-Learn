import express from "express";
import mysql from "mysql2/promise"

const app = express()
const PORT = process.env.PORT || 3000
//middleware
app.use(express.json())

import { Connector } from "@google-cloud/cloud-sql-connector";
//cloud sql connector 
const connector = new Connector();

async function getDbConfig() {
    const clientOps = await connector.getOptions({
        instanceConnectionName: process.env.INSTANCE_CONNECTION_NAME,
        authType: 'PASSWORD',
    })
    return {
        ...clientOps,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME,
    }
}

app.get('/', (req, res) => {
    res.json({
        status: 'success',
        message: 'Node.js Cloud SQL App is running',
        timestamp: new Date().toISOString()
    });
});

//health check
app.get('/health', (req, res) => {
    res.json({
        status: 'healthy',
        uptime: process.uptime(),
        timestamp: new Date().toISOString()
    });
});

// Database connection
app.get("/health/db", async (req, res) => {
    try {
        const dbConfig = await getDbConfig();
        const connection = await mysql.createConnection(dbConfig)

        // Simple query to test
        const [rows] = await connection.execute('SELECT NOW()'); // ✅ await add karo

        await connection.end();

        console.log('✅ Database connection successful');
        console.table(rows);

        res.json({
            status: 'success',
            message: 'Database connection successful',
            data: rows[0]
        });
    } catch (error) {
        console.error('❌ Database connection failed:', error.message);
        console.error(error.code)

        console.error(error.sqlState)

        console.error(error.errno)
        res.status(500).json({
            status: 'error',
            message: 'Database connection failed',
            error: error.message
        });
    }
});


// Start server
app.listen(PORT, () => {
    console.log('='.repeat(50));
    console.log(`🚀 Server is running on PORT: ${PORT}`);
    console.log(`📦 Instance: ${process.env.INSTANCE_CONNECTION_NAME}`);
    console.log(`💾 Database: ${process.env.DB_NAME}`);
    console.log(`👤 User: ${process.env.DB_USER}`);
    console.log('='.repeat(50));
    console.log(`\n📍 Endpoints:`);
    console.log(`   GET  http://localhost:${PORT}/`);
    console.log(`   GET  http://localhost:${PORT}/health`);
    console.log(`   GET  http://localhost:${PORT}/health/db`);
    console.log('='.repeat(50));
});