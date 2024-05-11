import React from "react";
import "./ArrayTable.style.css";

export default function ArrayTable({ head, side, data }) {
  let thead = null;
  if (head) {
    thead = (
      <thead>
        <tr>
          {Object.keys(head).map((element) => (
            <th onClick={head[element]} key={element} className="table-top-row">
              {element}
            </th>
          ))}
        </tr>
      </thead>
    );
  }

  return (
    <table className="main-table">
      {thead}
      <tbody>
        {data.map((row, rowIndex) => (
          <tr key={rowIndex}>
            {/* Conditional rendering for side column */}
            {side &&
              side[rowIndex] &&
              (rowIndex === 0 ? (
                <th key={`side_${rowIndex}`} className="table-left-column">
                  {side[rowIndex]}
                </th>
              ) : (
                <td key={`side_${rowIndex}`} className="table-left-column">
                  {side[rowIndex]}
                </td>
              ))}
            {/* Conditional rendering for rounded corners */}
            {!head && rowIndex === 0 && (
              <>
                {Object.keys(row).map((key) => (
                  <th key={key}>{row[key]}</th>
                ))}
              </>
            )}
            {/* Regular cells */}
            {!(!head && rowIndex === 0) &&
              Object.keys(row).map((key) => <td key={key}>{row[key]}</td>)}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
