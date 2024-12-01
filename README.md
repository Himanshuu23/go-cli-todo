    <h1>CLI Todo Application</h1>
    <p>A simple command-line todo application built using GoLang. Below are the commands and their usage:</p>

    <img src="path/to/your/image.png" alt="CLI Todo App Screenshot" style="width:100%; max-width:600px; margin-bottom:20px;">

    <h2>Commands</h2>
    <table border="1" cellspacing="0" cellpadding="10">
        <thead>
            <tr>
                <th>Command</th>
                <th>Description</th>
                <th>Example</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td><code>-help</code></td>
                <td>Displays all the available commands.</td>
                <td><code>./todo -help</code></td>
            </tr>
            <tr>
                <td><code>-list</code></td>
                <td>Prints all the todos in a table format.</td>
                <td><code>./todo -list</code></td>
            </tr>
            <tr>
                <td><code>-add "todo_title"</code></td>
                <td>Adds a new todo with the specified title.</td>
                <td><code>./todo -add "Buy groceries"</code></td>
            </tr>
            <tr>
                <td><code>-toggle index</code></td>
                <td>Toggles the completion status of the todo at the specified index.</td>
                <td><code>./todo -toggle 2</code></td>
            </tr>
            <tr>
                <td><code>-delete index</code></td>
                <td>Deletes the todo at the specified index.</td>
                <td><code>./todo -delete 3</code></td>
            </tr>
            <tr>
                <td><code>-edit index "new_title"</code></td>
                <td>Edits the title of the todo at the specified index.</td>
                <td><code>./todo -edit 1 "Go for a walk"</code></td>
            </tr>
        </tbody>
    </table>
