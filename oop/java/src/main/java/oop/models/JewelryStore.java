package oop.models;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
* Class JewelryStore represents a jewelry store.
*/
public class JewelryStore {
    private List<Jewelry> jewelries;
    private List<Customer> customers;
    private List<Sale> sales;

    private ObjectMapper mapper = new ObjectMapper();

    /** 
    * Class constructor.
    */
    public JewelryStore() {
    }

    /** 
    * Class constructor with parameters.
    */
    public JewelryStore(List<Jewelry> jewelries, List<Customer> customers, List<Sale> sales) {
        setJewelries(jewelries);
        setCustomers(customers);
        setSales(sales);
    }

    /** 
    * Class constructor from JSON file.
    */
    public JewelryStore(String fileName) throws IOException {
        String json = Files.readString(Paths.get(fileName));
        JewelryStore store = new JewelryStore();
        try {
            store = mapper.readValue(json, JewelryStore.class);
        } catch (JsonProcessingException e) {
            System.err.println(e.getMessage());
        }

        setJewelries(store.getJewelries());
        setCustomers(store.getCustomers());
        setSales(store.getSales());
    }

    public void setJewelries(List<Jewelry> jewelries) {
        this.jewelries = jewelries;
    }

    public List<Jewelry> getJewelries() {
        return jewelries;
    }

    public void setCustomers(List<Customer> customers) {
        this.customers = customers;
    }

    public List<Customer> getCustomers() {
        return customers;
    }

    public void setSales(List<Sale> sales) {
        this.sales = sales;
    }

    public List<Sale> getSales() {
        return sales;
    }

    public void addJewelry(Jewelry jewelry) {
        jewelries.add(jewelry);
    }

    public void removeJewelry(Jewelry jewelry) {
        jewelries.remove(jewelry);
    }

    public void addCustomer(Customer customer) {
        customers.add(customer);
    }

    public void removeCustomer(Customer customer) {
        customers.remove(customer);
    }

    public void addSale(Sale sale) {
        sales.add(sale);
    }

    public void removeSale(Sale sale) {
        sales.remove(sale);
    }

    /**
    * toJSON makes JSON string from Jewelry store instance.
    */
    private String toJSON() throws JsonProcessingException {
        return mapper.writeValueAsString(this);
    }

    /**
    * writeJSON writes Jewelry store as JSON to filename param.
    *
    * @param filename name of file.
    */
    public void writeJSON(String fileName) throws IOException {
        try {
            String json = toJSON();
            Path filePath = Paths.get(fileName);
            Files.writeString(filePath, json);
        } catch (JsonProcessingException e) {
            System.err.println(e.getMessage());
        }
    }

    /**
    * writeJewelriesToDB writes jewelries of store to DB using connection
    *
    * @param connection connection to DB.
    */
    public void writeJewelriesToDB(Connection connection) {
        for (Jewelry jewelry : jewelries) {
            jewelry.saveToDB(connection);
        }
    }

    /**
    * readJewelriesToDB reads jewelries of store from DB using connection
    *
    * @param connection connection to DB.
    */
    public void readJewelriesFromDB(Connection connection) {
        try {
            Statement statement = connection.createStatement();
            statement.setQueryTimeout(30);
            ResultSet rs = statement.executeQuery("SELECT id, name, jewelry_cost, making_cost FROM jewelries");
            List<Jewelry> jewelries = new ArrayList<Jewelry>();
            while (rs.next()) {
                jewelries.add(new Jewelry(rs.getInt("id"), rs.getString("name"), rs.getInt("jewelry_cost"),
                        rs.getInt("making_cost")));
            }

            setJewelries(jewelries);
        } catch (SQLException e) {
            System.err.println("Error reading jewelries: " + e.getMessage());
        }
    }

    /**
    * writeCustomersToDB writes customers of store to DB using connection
    *
    * @param connection connection to DB.
    */
    public void writeCustomersToDB(Connection connection) {
        for (Customer customer : customers) {
            customer.saveToDB(connection);
        }
    }

    /**
    * readCustomersFromDB reads customers of store from DB using connection
    *
    * @param connection connection to DB.
    */
    public void readCustomersFromDB(Connection connection) {
        try {
            Statement statement = connection.createStatement();
            statement.setQueryTimeout(30);
            ResultSet rs = statement.executeQuery("SELECT id, name, surname, phone, address FROM customers");
            List<Customer> customers = new ArrayList<Customer>();
            while (rs.next()) {
                customers.add(new Customer(rs.getInt("id"), rs.getString("name"), rs.getString("surname"),
                        rs.getString("phone"), rs.getString("address")));
            }

            setCustomers(customers);
        } catch (SQLException e) {
            System.err.println("Error reading customers: " + e.getMessage());
        }
    }

    /**
    * writeSalesToDB writes sales of store to DB using connection
    *
    * @param connection connection to DB.
    */
    public void writeSalesToDB(Connection connection) {
        for (Sale sale : sales) {
            sale.saveToDB(connection);
        }
    }

    /**
    * readSalesFromDB reads sales of store from DB using connection
    *
    * @param connection connection to DB.
    */
    public void readSalesFromDB(Connection connection) {
        try {
            Statement statement = connection.createStatement();
            statement.setQueryTimeout(30);
            ResultSet rs = statement.executeQuery("SELECT id, jewelry_id, customer_id FROM sales");
            List<Sale> sales = new ArrayList<Sale>();
            while (rs.next()) {
                Statement jStmt = connection.createStatement();
                jStmt.setQueryTimeout(30);
                ResultSet js = jStmt
                        .executeQuery("SELECT id, name, jewelry_cost, making_cost FROM jewelries WHERE id == "
                                + rs.getInt("jewelry_id"));
                Jewelry j = new Jewelry();
                while (js.next()) {
                    j = new Jewelry(js.getInt("id"), js.getString("name"), js.getInt("jewelry_cost"),
                            js.getInt("making_cost"));
                }

                Statement cStmt = connection.createStatement();
                cStmt.setQueryTimeout(30);
                ResultSet cs = cStmt.executeQuery("SELECT id, name, surname, phone, address FROM customers WHERE id == "
                        + rs.getInt("customer_id"));
                Customer c = new Customer();
                while (cs.next()) {
                    c = new Customer(cs.getInt("id"), cs.getString("name"), cs.getString("surname"),
                            cs.getString("phone"), cs.getString("address"));
                }

                sales.add(new Sale(rs.getInt("id"), j, c));
            }

            setSales(sales);
        } catch (SQLException e) {
            System.err.println("Error reading sales: " + e.getMessage());
        }
    }
}
