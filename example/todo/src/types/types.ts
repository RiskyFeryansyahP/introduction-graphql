export type Todo = {
    id: string;
    text: string;
    done: boolean;
}

export type Todos = {
    todos: Array<Todo>
}