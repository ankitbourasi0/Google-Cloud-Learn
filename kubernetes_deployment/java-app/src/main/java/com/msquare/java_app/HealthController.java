package com.msquare.java_app;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.sql.DataSource;
import java.sql.Connection;

@RestController
public class HealthController {

    @Autowired
    private DataSource dataSource;
    @GetMapping("/health/db")
    public String checkDatabase(){
        try (Connection conn = dataSource.getConnection()) {
            return "Welcome to Cloud SQL, Database connected successfully!";
        } catch (Exception e) {
            return "Database connection failed: " + e.getMessage();
        }
    }
}
