import React, { useEffect, useState } from "react";
import "./UserControl.style.css";
import "../../UI/GreyContainer/GreyContainer.style.css";
import MagnifyingGlass from "../../Images/MagnifyingGlass.svg";
import { ArrayTable, UserCard } from "../../UI";
import { GetUsers } from "../../Services/Users";

export default function UserControl() {
  const head = {
    ID: function () {
      console.log("1");
    },
    Username: function () {
      console.log("2");
    },
    Email: function () {
      console.log("3");
    },
    Credits: function () {
      console.log("4");
    },
    "Is blocked": function () {
      console.log("5");
    },
    Role: function () {
      console.log("6");
    },
    "Creation date": function () {
      console.log("7");
    },
  };

  const [users, setUsers] = useState([]);
  useEffect(() => {
    (async function () {
      const data = (await GetUsers())?.Users;
      setUsers(data);
    })();
  }, []);

  users.forEach((u) => {
    delete u["ProfilePicture"];
    delete u["UpdatedAt"];
    u["IsBlocked"] = u["IsBlocked"] ? "No" : "Yes";
  });

  const [selectedUser, setSelectedUser] = useState(null);

  const clickHandler = (index) => {
    return () => {
      setSelectedUser(users[index]);
    };
  };

  return (
    <div className="user-control-page">
      <main className="main">
        <div className="main-content">
          <div className="main-content-block-search">
            <div className="main-content-block-searchbar">
              <img src={MagnifyingGlass} alt="magnifyingGlass" />
              <input
                className="header-search"
                type="text"
                placeholder="Find user..."
              />
              <button type="button" className="header-search-button">
                ➜
              </button>
            </div>
          </div>
          <div className="main-content-block-main">
            <div className="main-content-block">
              <div className="main-description-table grey-container">
                <ArrayTable
                  data={users}
                  // side={side}
                  head={head}
                  className="main-table"
                  clickHandler={clickHandler}
                />
              </div>
              {/* <div className="main-description-table grey-container">
                <ObjectTable data={obj} />
              </div> */}
            </div>
            {selectedUser ? (
              <UserCard
                username={selectedUser.EpicID}
                email={selectedUser.Email}
                credits={selectedUser.Balance}
                date={selectedUser.CreatedAt}
              />
            ) : null}
          </div>
        </div>
      </main>
    </div>
  );
}
