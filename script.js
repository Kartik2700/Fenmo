document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('expenseForm');
    const filterCategory = document.getElementById('filterCategory');
    const tbody = document.getElementById('expenseTableBody');
    const totalDisplay = document.getElementById('totalDisplay');

    // 1. Initial Load
    loadExpenses();

    // 2. Event Listener for Form Submission
    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        const payload = {
            amount: parseInt(document.getElementById('rupees').value) || 0,
            category: document.getElementById('category').value,
            description: document.getElementById('description').value,
            date: document.getElementById('date').value
        };

        try {
            const response = await fetch('http://localhost:8080/expenses', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            if (response.ok) {
                form.reset();
                loadExpenses(); // Refresh list after adding
            } else {
                alert("Failed to save expense.");
            }
        } catch (error) {
            console.error('Error:', error);
        }
    });

    // 3. Event Listener for Filter Dropdown
    filterCategory.addEventListener('change', loadExpenses);

    // 4. Function to Fetch and Display Data
    async function loadExpenses() {
        const category = filterCategory.value;
        const url = category ? `http://localhost:8080/expenses?category=${category}` : 'http://localhost:8080/expenses';

        try {
            const res = await fetch(url);
            const data = await res.json();

            renderTable(data);
        } catch (error) {
            console.error('Error loading expenses:', error);
        }
    }

    // 5. Helper Function to build the table and sum the total
    function renderTable(expenses) {
        tbody.innerHTML = '';
        let totalAmount = 0;

        expenses.forEach(exp => {
            totalAmount += exp.amount;

            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${exp.date}</td>
                <td>${exp.category}</td>
                <td>${exp.description}</td>
                <td>₹${exp.amount}</td>
            `;
            tbody.appendChild(row);
        });

        // Update Total display
        totalDisplay.innerText = `₹${(totalAmount)}`;
    }
});