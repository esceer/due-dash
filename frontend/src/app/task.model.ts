export interface Task {
  id: number;
  title: string;
  status: 'active' | 'completed';
  category: string;
  due_date?: string;
}