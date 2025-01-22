import axios from "axios";
import { Course } from "../model/canvas/course";

export async function listCourses() {
    const resp = await axios.get<Course[]>("/api/courses");
    return resp.data;
}