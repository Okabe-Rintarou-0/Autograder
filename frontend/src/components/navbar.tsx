import { Col, Row } from "antd";
import React from "react";
import { Link } from "react-router-dom";

export default function NavBar() {
    return (
        <Row className="navbar" justify="start">
            <Col>
                <Link to="/">Book Store</Link>
            </Col>
            <Col flex="auto">
            </Col>
        </Row>
    );
}