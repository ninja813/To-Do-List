# Full-Stack Todo App

A modern todo application built with **React** frontend and **Go** backend.

## 🚀 Features

- ✅ Add new tasks
- ✅ Mark tasks as complete/incomplete
- ✅ Edit existing tasks
- ✅ Delete tasks
- ✅ Real-time updates
- ✅ Responsive design
- ✅ Modern UI with gradients and animations

## 🏗️ Architecture

### Backend (Go)
- **HTTP Server**: Built-in Go HTTP server
- **REST API**: Full CRUD operations
- **CORS Support**: Cross-origin requests enabled
- **JSON**: Data exchange format
- **In-memory Storage**: Tasks stored in memory (resets on restart)

### Frontend (React)
- **React 18**: Modern React with hooks
- **Components**: Modular component architecture
- **API Integration**: Fetch API for backend communication
- **Responsive Design**: Mobile-friendly interface
- **Modern CSS**: Gradients, animations, and clean design

## 📁 Project Structure

```
├── main.go                 # Go backend server
├── frontend/               # React frontend
│   ├── public/
│   │   └── index.html
│   ├── src/
│   │   ├── components/
│   │   │   ├── AddTask.js
│   │   │   ├── TaskItem.js
│   │   │   └── TaskList.js
│   │   ├── App.js
│   │   ├── index.js
│   │   └── index.css
│   └── package.json
└── README.md
```

## 🛠️ Prerequisites

- **Go 1.19+**: [Download Go](https://golang.org/dl/)
- **Node.js 16+**: [Download Node.js](https://nodejs.org/)
- **npm**: Comes with Node.js

## 🚀 How to Run

### Option 1: Development Mode (Recommended)

1. **Start the Go backend**:
   ```bash
   go run main.go
   ```
   The backend will start on `http://localhost:8080`

2. **In a new terminal, start the React frontend**:
   ```bash
   cd frontend
   npm install
   npm start
   ```
   The frontend will start on `http://localhost:3000`

3. **Open your browser** and go to `http://localhost:3000`

### Option 2: Production Mode

1. **Build the React app**:
   ```bash
   cd frontend
   npm install
   npm run build
   ```

2. **Start the Go server**:
   ```bash
   go run main.go
   ```

3. **Open your browser** and go to `http://localhost:8080`

## 🔧 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tasks` | Get all tasks |
| POST | `/api/tasks` | Create a new task |
| PUT | `/api/tasks/{id}` | Update a task |
| DELETE | `/api/tasks/{id}` | Delete a task |

### Example API Usage

```bash
# Get all tasks
curl http://localhost:8080/api/tasks

# Create a new task
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "completed": false}'

# Update a task
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "completed": true}'

# Delete a task
curl -X DELETE http://localhost:8080/api/tasks/1
```

## 🎨 Features Explained

### Backend Features
- **CORS Middleware**: Enables cross-origin requests
- **JSON Handling**: Automatic JSON encoding/decoding
- **Error Handling**: Proper HTTP status codes
- **RESTful Design**: Standard REST API patterns

### Frontend Features
- **React Hooks**: useState, useEffect for state management
- **Component Architecture**: Modular, reusable components
- **Real-time Updates**: Immediate UI updates after API calls
- **Error Handling**: User-friendly error messages
- **Responsive Design**: Works on desktop and mobile

## 🎯 Key Go Concepts Demonstrated

- **HTTP Server**: `net/http` package
- **JSON Marshaling**: `encoding/json` package
- **Structs**: Custom data types with JSON tags
- **Methods**: Functions attached to types
- **Error Handling**: Proper error management
- **CORS**: Cross-origin resource sharing
- **File Serving**: Static file hosting

## 🎯 Key React Concepts Demonstrated

- **Functional Components**: Modern React with hooks
- **State Management**: useState and useEffect
- **Props**: Data passing between components
- **Event Handling**: User interactions
- **API Integration**: Fetch API usage
- **Conditional Rendering**: Dynamic UI updates

## 🔄 Development Workflow

1. **Backend Changes**: Restart Go server (`go run main.go`)
2. **Frontend Changes**: React hot-reloads automatically
3. **API Testing**: Use browser dev tools or curl commands
4. **Debugging**: Check browser console and Go server logs

## 🚀 Deployment Options

### Local Development
- Backend: `http://localhost:8080`
- Frontend: `http://localhost:3000`

### Production
- Build React app: `npm run build`
- Serve from Go: `http://localhost:8080`

## 🎉 What You've Built

A complete full-stack application with:
- ✅ Modern React frontend
- ✅ Go backend with REST API
- ✅ Real-time CRUD operations
- ✅ Beautiful, responsive UI
- ✅ Proper error handling
- ✅ CORS support for development

This demonstrates the power of combining Go's simplicity with React's component-based architecture for building modern web applications!
