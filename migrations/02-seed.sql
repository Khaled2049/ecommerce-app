-- Insert sample customers
INSERT INTO customers (name, email, phone, preferences) VALUES
('John Doe', 'john.doe@example.com', '123-456-7890', '{"newsletter": true, "notifications": false}'),
('Jane Smith', 'jane.smith@example.com', '987-654-3210', '{"newsletter": false, "notifications": true}'),
('David Lee', 'david.lee@example.com', '555-123-4567', NULL),
('Sarah Jones', 'sarah.jones@example.com', '111-222-3333', '{"newsletter": true}');

-- Insert sample categories (with parent-child relationships)
INSERT INTO categories (name, description, parent_category_id) VALUES
('Electronics', 'Electronic devices', NULL),
('Computers', 'Desktop and laptop computers', 1), -- Parent: Electronics
('Laptops', 'Portable computers', 2),       -- Parent: Computers
('Phones', 'Mobile phones', 1),           -- Parent: Electronics
('Clothing', 'Apparel and accessories', NULL),
('Menswear', 'Clothing for men', 5);

-- Insert sample products
INSERT INTO products (name, description, price, stock, tags) VALUES
('Laptop X1', 'High-performance laptop', 1200.00, 50, ARRAY['laptop', 'computer', 'electronics']),
('Smartphone Z2', 'Latest smartphone model', 800.00, 100, ARRAY['phone', 'electronics', 'new']),
('T-Shirt Basic', 'Cotton t-shirt', 20.00, 200, ARRAY['clothing', 'casual', 'mens']),
('Office Chair Ergonomic', 'Ergonomic office chair', 300.00, 30, ARRAY['furniture', 'office']),
('Laptop Y2', 'Budget friendly laptop', 500.00, 75, ARRAY['laptop', 'computer', 'electronics']);

-- Insert sample product_categories mappings
INSERT INTO product_categories (product_id, category_id) VALUES
(1, 3), -- Laptop X1 is in Laptops
(2, 4), -- Smartphone Z2 is in Phones
(3, 6), -- T-Shirt Basic is in Menswear
(1,2),
(5,2),
(5,3);

-- Insert sample payment methods
INSERT INTO payment_methods (customer_id, type, provider, account_number, expiry_date, is_default) VALUES
(1, 'Credit Card', 'Visa', '1111222233334444', '2025-12-31', true),
(2, 'PayPal', 'PayPal', 'jane.smith@paypal.com', NULL, true),
(1, 'Debit Card','MasterCard', '5555666677778888','2024-03-15',false);

-- Insert sample orders
INSERT INTO orders (customer_id, total_amount, status) VALUES
(1, 1200.00, 'completed'),
(2, 800.00, 'processing'),
(1, 320.00, 'completed');

-- Insert sample order items
INSERT INTO order_items (order_id, product_id, quantity, unit_price, subtotal) VALUES
(1, 1, 1, 1200.00, 1200.00),
(2, 2, 1, 800.00, 800.00),
(3, 3, 2, 20.00, 40.00),
(3, 4,1,300,300);

-- Insert sample payments
INSERT INTO payments (order_id, payment_method_id, amount, status, transaction_id) VALUES
(1, 1, 1200.00, 'completed', 'TXN12345'),
(2, 2, 800.00, 'processing', 'TXN67890'),
(3,1,320,'completed','TXN98765');