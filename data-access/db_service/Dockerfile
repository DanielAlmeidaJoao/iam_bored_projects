FROM mysql:latest

# Set environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=secret
ENV MYSQL_DATABASE=recordings 
ENV MYSQL_USER=myuser 
ENV MYSQL_PASSWORD=secret

# Copy the SQL script to the container, which MySQL automatically executes on first startup
COPY create-tables.sql /docker-entrypoint-initdb.d/

# Expose MySQL port
EXPOSE 3306
