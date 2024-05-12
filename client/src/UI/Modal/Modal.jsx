import React from "react";
import LogoMain from "../../Images/LogoMain.svg"

export default function Modal({ children, onClose }) {
  return (
    <div className="modal-backdrop" onClick={onClose}>
      <div className="modal-content"> {/*onClick={(e) => e.stopPropagation()} */}
        <div className="header">
          <img src={LogoMain} alt="Logo" width="100%" height="30" />
        </div>
        <div className="body">
          {children}
        </div>
      </div>
    </div>
  )
}