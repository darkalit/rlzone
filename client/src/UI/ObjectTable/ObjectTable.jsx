import React from "react";
import "./ObjectTable.style.css";

export default function ObjectTable({ data }) {
  for (let e of Object.keys(data)) {
    if (typeof data[e] === typeof true) {
      data[e] = data[e] ? "Yes" : "No";
    }
  }

  const arr = Object.entries(data);

  return (
    <table className="main-table">
      <tbody>
        {arr.map((row, rowIndex) => (
          <tr key={rowIndex}>
            <td>{row[0]}</td>
            <td>{row[1]}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
