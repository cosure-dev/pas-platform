# Stage 1: Build stage with dependencies
FROM python:3.11-slim-bookworm AS builder

# Set the working directory
WORKDIR /usr/src/app

# Set environment variables for Python
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Install dependencies
COPY requirements.txt .
RUN pip wheel --no-cache-dir --no-deps --wheel-dir /usr/src/app/wheels -r requirements.txt

# Stage 2: Final production image
FROM python:3.11-slim-bookworm

# Create a non-root user for security
RUN addgroup --system app && adduser --system --group app

# Set working directory
WORKDIR /home/app

# Copy the installed wheels from the builder stage
COPY --from=builder /usr/src/app/wheels /wheels

# Install the wheels
RUN pip install --no-cache /wheels/*

# Copy the application source code
COPY . .

# Change ownership to the non-root user
RUN chown -R app:app /home/app

# Switch to the non-root user
USER app

# Expose the port the FastAPI service will run on
EXPOSE 8000

# Command to run the application using uvicorn
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]