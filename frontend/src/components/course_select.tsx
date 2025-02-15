import { Select } from "antd";
import { Course } from "../model/canvas/course";

export interface CourseSelectProps {
    courses: Course[];
    disabled?: boolean;
    onChange?: (courseId: number) => void;
    value?: number;
}

export default function CourseSelect({ courses, disabled, onChange, value }: CourseSelectProps) {
    const formatCourses = (courses: Course[]) => {
        const formatted: Course[] = [];
        courses.map(course => {
            const term = course.term.name.replace("Spring", "春").replace("Fall", "秋");
            formatted.push({
                ...course,
                name: `${course.name}(${term}, ${course.teachers?.[0]?.display_name ?? '未知教师'})`
            });
        });
        // sort by term id, latest first
        formatted.sort((a, b) => b.term.id - a.term.id);
        return formatted;
    }

    const courseLabel = (course: Course) => {
        return course.enrollments.find(enrollment => enrollment.role === "TaEnrollment") ?
            <span><span style={{ color: "red" }}>*</span>{course.name}</span> :
            course.name
    }

    let formattedCourses = formatCourses(courses);

    return <Select
        showSearch
        optionFilterProp="children"
        filterOption={(filter, options) => {
            if (!options || !filter) {
                return true;
            }
            let courseId = options.value;
            let course = courses.find(course => course.id === courseId)!;
            // filter by course name or teacher name
            return course.name.includes(filter) || course.teachers.filter(teacher => teacher.display_name.includes(filter)).length > 0;
        }}
        value={value}
        style={{ width: 350 }}
        disabled={disabled}
        onChange={onChange}
        placeholder={"请选择课程"}
        options={formattedCourses.map(course => ({
            label: courseLabel(course),
            value: course.id
        }))}
    />
}