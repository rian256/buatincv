#!/usr/bin/expect

spawn bin/main new

expect "Enter your name\r"

send -- "Fulana\r"

expect "Enter your role\r"

send -- "Graphic Designer Specialist\r"

expect "Enter your current address\r"

send -- "Banten, Indonesia | Phone: +623456789 | Email: rosellia@test.com\r"

expect "Enter your summary"

send -- "Passionate graphic designer with over 3 years of experience with branding projects, with past clients that include a Fortune 20 company. Possesses a strong understanding of vector art and effective design, which allows the prompt delivery of clean and memorable logo designs for clients.\r"

expect "Enter the name and location of the company where you previously worked or currently work."

send -- "PT PUH SEPUH FOUNDATION - Bandung, Indonesia\r"

expect "Enter your experience and indicate the start and end dates of your employment.:"

send -- "Graphic Designer (2018-2022)\r"

expect "Describe your job responsibilities\r"

send -- "- Design graphics to meet specific promotional and commercial needs like logos, packaging, displays, or imagery for digital and printable products\r"

expect  "Would you like to add more information about the job responsibilities? y/n"

send -- "y\r"

expect "Describe your job responsibilities\r"

send -- "- Decide how images and text work to fit a specific layout\r"

expect "Would you like to add more information about the job responsibilities? y/n"

send -- "y\r"

expect "Describe your job responsibilities\r"

send -- "- Develop illustrations, logos and other designs using software or by hand.\r"

expect "Would you like to add more information about the job responsibilities? y/n"

send -- "y\r"

expect "Describe your job responsibilities\r"

send -- "- Conceptualize visuals based on requirements.\r"

expect "Would you like to add more information about the job responsibilities? y/n"

send -- "n\r"

expect "Enter your formal education experience\r"

send -- "XYZ University - (2014-2018)\r"

expect "Enter your Educational Degree"

send -- "Bachelor of Design - GPA 4.0\r"

expect "Enter your skill:"

send -- "- Figma\r"

expect "Do you want to add more skills? y/n"

send -- "y\r"

expect "Enter your skill:"

send -- "- Adobe XD\r"

expect "Do you want to add more skills? y/n"

send -- "y\r" 

expect "Enter your skill:"

send -- "- Adobe After Effects\r"

expect "Do you want to add more skills? y/n"

send -- "y\r"

expect "Enter your skill:"

send -- "- Adobe Photoshop\r"

expect "Do you want to add more skills? y/n"

send -- "y\r" 

expect "Enter your skill:"

send -- "- Adobe Illustrator\r"

expect "Do you want to add more skills? y/n"

send -- "n\r" 

expect eof