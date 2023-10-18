CREATE TABLE Heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255) NOT NULL,
    Skill VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL
);

CREATE TABLE Villains (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Universe VARCHAR(100) NOT NULL,
    Skill VARCHAR(100) NOT NULL,
    ImageURL VARCHAR(100) NOT NULL
);

CREATE TABLE Inventories (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    ItemCode VARCHAR(100),
    Stock INT,
    Description VARCHAR(100),
    Status VARCHAR(100)
);

CREATE TABLE CriminalReports (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    hero_ID INT,
    villain_ID INT,
    description VARCHAR(100),
    time DATETIME,
    FOREIGN KEY (hero_ID) REFERENCES Heroes(ID),
    FOREIGN KEY (villain_ID) REFERENCES Villains(ID)
);


INSERT INTO Heroes (Name, Universe, Skill, ImageURL)
VALUES
    ('ABADDON', 'UNIVERSAL', 'SHIELDS HIS ALLIES OR HIMSELF FROM ATTACKS', 'ABADDON.jpg'),
    ('ANTI-MAGE', 'AGILITY', 'SLASHES HIS FOES WITH MANA-DRAINING ATTACKS', 'ANTI-MAGE.jpg'),
    ('ANCIENT APPARITION', 'INTELLIGENCE', 'LAUNCHES A POWERFUL ICY BLAST FROM ANYWHERE ON THE MAP', 'ANCIENT APPARITION.jpg');

INSERT INTO Villains (Name, Universe, Skill, ImageURL)
VALUES
    ('DEATH PROPHET', 'INTELLIGENCE', 'SUMMONS AN ARMY OF GHOSTS TO ATTACK' , 'PROPHET.jpg'),
    ('LYCAN', 'UNIVERSAL', 'SHAPESHIFTS INTO A FEROCIOUS BEAST' , 'LYCAN.jpg'),
    ('NIGHT STALKER', 'STRENGTH', 'ENPOWERED BY THE SHADOWS OF NIGHTFALL' , 'STALKER.jpg');

INSERT INTO Inventories (Name, ItemCode, Stock, Description, Status)
VALUES 
    ('Crystalys', 'ITM001', 10, 'A blade forged from rare crystals, it seeks weak points in enemy armor.', 'active'),
    ('Daedalus', 'ITM002', 5, 'A weapon of incredible power that is difficult for even the strongest of warriors to control.', 'active'),
    ('Butterfly', 'ITM003', 8, 'Only the mightiest and most experienced of warriors can wield the Butterfly, but it provides incredible dexterity in combat.', 'broken'),
    ('Eye of Skadi', 'ITM004', 12, 'Extremely rare artifact, guarded by the azure dragons.', 'active'),
    ('Yasha and Kaya', 'ITM005', 3, 'Yasha and Kaya when paired together share a natural resonance.', 'broken');

INSERT INTO CriminalReports (hero_ID, villain_ID, description, time)
VALUES 
    (1, 1, 'Burglary at Main Street', '2023-10-17 18:30:00'),
    (2, 2, 'Bank Robbery on Elm Avenue', '2023-10-16 14:45:00'),
    (3, 3, 'Kidnapping Incident at Park Square', '2023-10-15 09:20:00'),
    (1, 3, 'Arson in Downtown', '2023-10-14 20:10:00');


