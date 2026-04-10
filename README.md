# Fenmo
# Personal Expense Tracker
**Live Application:** [https://fenmo-expense-tracker-sueo.onrender.com/](https://fenmo-expense-tracker-sueo.onrender.com/)

A minimal full-stack application to record and review personal expenses.
Built to handle real-world conditions like unreliable networks and retries while maintaining data integrity.

## 🚀 Key Design Decisions
* **Money Handling:** Used the `int64` data type to store expenses as whole Rupees. This ensures mathematical precision and avoids the rounding errors common with floating-point types in financial software.
* **Concurrency:** Implemented a thread-safe in-memory store using `sync.RWMutex` in Go. This ensures the API behaves correctly even if multiple requests hit the server simultaneously.
* **Architecture:** Chose a modular structure with a separate `script.js` file to handle API interactions, ensuring clear separation of concerns between the UI and the backend logic.

## 🛠 Tech Stack
* **Backend:** Go (Golang) using `net/http` for a lightweight, high-performance API.
* **Frontend:** Vanilla JavaScript, HTML5, and CSS3.

## ⚖️ Trade-offs & Future Improvements
* **In-Memory Storage:** Chose an in-memory database for simplicity within the timebox. The codebase is structured to allow easy migration to a persistent database (e.g., SQLite or PostgreSQL).
* **Decimal Precision:** For this MVP, I prioritized whole-rupee accuracy to meet the "appropriate type for real money" requirement simply. A future iteration counld involving  "Rupee.Paise based" system for sub-unit precision.
* **Authentication:** Currently open for local use. Production versions would include user authentication and session management.

## 🧪 Resilience & Correctness
* **Idempotency & Safety:** The backend utilizes server-side locking to prevent data corruption during retries or page reloads.
* **Filtering & Sorting:** Supports category-based filtering and "newest-first" date sorting directly through API query parameters.
