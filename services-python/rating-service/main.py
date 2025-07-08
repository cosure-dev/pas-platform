from fastapi import FastAPI

# Initialize a new FastAPI application instance
# The 'uvicorn' command in our Dockerfile looks for this 'app' object
app = FastAPI()

# Define a simple health check endpoint
@app.get("/health")
def read_root():
    """
    This endpoint confirms the service is running.
    """
    return {"status": "UP", "service": "rating-service"}