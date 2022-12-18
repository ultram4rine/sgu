package oop.models;

import java.sql.Connection;
import java.sql.SQLException;
import java.sql.Statement;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
* Class Sale represents a sale in pawnshop.
*/
public class Sale {
    private int id;
    private Jewelry jewelry;
    private Customer customer;

    /** 
    * Class constructor.
    */
    public Sale() {
    }

    /** 
    * Class constructor with parameters.
    */
    public Sale(int id, Jewelry jewelry, Customer customer) {
        setId(id);
        setJewelry(jewelry);
        setCustomer(customer);
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getId() {
        return id;
    }

    public void setJewelry(Jewelry jewelry) {
        this.jewelry = jewelry;
    }

    public Jewelry getJewelry() {
        return jewelry;
    }

    public void setCustomer(Customer customer) {
        this.customer = customer;
    }

    public Customer getCustomer() {
        return customer;
    }

    /**
    * toJSON makes JSON string from Sale instance.
    */
    public String toJSON() throws JsonProcessingException {
        return new ObjectMapper().writeValueAsString(this);
    }

    /**
    * saveToDB saves Sale instance to DB using connection.
    * 
    * @param connection connection to DB.
    */
    public void saveToDB(Connection connection) {
        try {
            Statement statement = connection.createStatement();
            statement.setQueryTimeout(30);
            statement.executeUpdate("INSERT INTO sales(id, jewelry_id, customer_id) VALUES (" + getId() + ", "
                    + getJewelry().getId() + ", " + getCustomer().getId() + ")");
        } catch (SQLException e) {
            System.err.println("Error inserting sale: " + e.getMessage());
        }
    }
}
