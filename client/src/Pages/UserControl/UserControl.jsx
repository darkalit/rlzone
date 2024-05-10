import React from "react";
import "./UserControl.style.css";

export default function UserControl() {
  const styles = {
    background: "red",
    height: "100px",
    width: "100px",
  };

  return (
    <>
      <div style={styles}></div>
    </>
  );
}
