
interface EmailProps {
    email: string;
}

export function Email(props: EmailProps) {
    const { email } = props;
    return <a href={`mailto:${email}`}>{email}</a>
}