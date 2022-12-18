package oop.models;

import java.sql.Connection;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.regex.Pattern;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
* Class Customer represents a customer in pawnshop.
*/
public class Customer {
    private int id;
    private String name;
    private String surname;
    private String phone;
    private String address;

    // special regex for phone.
    private Pattern phoneRegex = Pattern.compile("^(\\d{3}[- .]?){2}\\d{4}$");

    /** 
    * Class constructor.
    */
    public Customer() {
    }

    /** 
    * Class constructor with parameters.
    */
    public Customer(int id, String name, String surname, String phone, String address) {
        setId(id);
        setName(name);
        setSurname(surname);
        setPhone(phone);
        setAddress(address);
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getId() {
        return id;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setSurname(String surname) {
        this.surname = surname;
    }

    public String getSurname() {
        return surname;
    }

    public void setPhone(String phone) throws IllegalArgumentException {
        if (phoneRegex.matcher(phone).matches()) {
            this.phone = phone;
        } else {
            throw new IllegalArgumentException("Wrong phone number.");
        }
    }

    public String getPhone() {
        return phone;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getAddress() {
        return address;
    }

    /**
    * toJSON makes JSON string from Customer instance.
    */
    public String toJSON() throws JsonProcessingException {
        return new ObjectMapper().writeValueAsString(this);
    }

    /**
    * saveToDB saves Customer instance to DB using connection.
    * 
    * @param connection connection to DB.
    */
    public void saveToDB(Connection connection) {
        try {
            Statement statement = connection.createStatement();
            statement.setQueryTimeout(30);
            statement.executeUpdate("INSERT INTO customers(id, name, surname, phone, address) VALUES (" + getId()
                    + ", '" + getName() + "', '" + getSurname() + "', '" + getPhone() + "', '" + getAddress() + "')");
        } catch (SQLException e) {
            System.err.println("Error inserting customer: " + e.getMessage());
        }
    }
}
