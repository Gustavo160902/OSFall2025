
Arkhipenko, A. (2024). TaskScheduler for Arduino [Software]. GitHub. https://github.com/arkhipenko/TaskScheduler 
 This repository provides a cooperative scheduling library that allows multiple timed tasks on Arduino boards. 
Difference: Our project uses this as the foundation but focuses on measuring its fairness, latency, and performance instead of extending the code show. 

 

Arduino Forum. (2024, January 15). TaskScheduler library – User experiences [Online forum post].  

https://forum.arduino.cc/t/taskscheduler-library-what-are-your-experiences/675944 
 This forum thread collects user experiences, debugging advice, real problems and solutions, and performance reports from the TaskScheduler library. 
Difference: We perform controlled experiments and quantitative measurements instead of anecdotal observations. 

 

Bacon, R. S. (2021, May 29). #253 Accurate task scheduler for the Arduino (and STM32, ESP32...) [Video]. YouTube. https://www.youtube.com/watch?v=eoJUlH_rWOE&t=670s 
 This video demonstrates building an accurate task scheduler for microcontrollers and visualizing its timing behavior. 
Difference: Our project uses the official TaskScheduler library rather than writing one from scratch and emphasizes fairness testing. 

 

Circuitstate. (2023, January 4). ptScheduler – A minimal cooperative task scheduler for Arduino. Circuitstate Electronics. https://www.circuitstate.com/tutorials/ptscheduler-a-minimal-cooperative-task-scheduler-for-arduino/ 
 This article presents a lightweight scheduler that runs multiple tasks using cooperative timing. 
Difference: We analyze an existing advanced library rather than designing a minimal new scheduler. 

 

FreeRTOS. (2024). FreeRTOS documentation. Real Time Engineers Ltd. https://www.freertos.org/Documentation/00-Overview 
 The FreeRTOS documentation explains preemptive scheduling, task priorities, and kernel timing used in embedded systems. 
Difference: Our study compares cooperative scheduling behavior against these preemptive methods for educational analysis. 

 

Kim, S., & Lee, J. (2021). Performance evaluation of lightweight scheduling in IoT devices. IEEE Internet of Things Journal, 8(5), 3892–3903. https://doi.org/10.1109/JIOT.2021.3051234 
 This research analyzes low-resource schedulers used in IoT microcontrollers. 
Difference: We reproduce a similar concept but limit our evaluation to Arduino TaskScheduler in cooperative mode. 

 

Liu, C. L., & Layland, J. W. (1973). Scheduling algorithms for multiprogramming in a hard-real-time environment. Journal of the ACM, 20(1), 46–61. https://doi.org/10.1145/321738.321743 
 This classic paper defines foundational real-time scheduling theory such as rate-monotonic and earliest-deadline-first algorithms. 
Difference: Our work uses these ideas conceptually but tests a practical cooperative implementation on Arduino hardware. 

 

Majumdar, S. (2022, August 10). Cooperative vs preemptive multitasking in embedded systems. Embedded.com.  

https://www.embedded.com/the-basics-of-embedded-multitasking-on-a-pic-part-1-reentrant-multitasking/?utm_source 
 Explains the trade-offs between cooperative and preemptive multitasking in embedded environments. 
Difference: We provide experimental data comparing these approaches rather than theoretical discussion. 

 

Moya, P. (2020, July 27). Arduino multitasking without RTOS. Hackster.io. 

https://www.hackster.io/onelife/multitasking-on-arduino-b5f216 
 Shows how to simulate multitasking on Arduino without an operating system. 
Difference: We rely on an actual scheduler library to test systematic fairness and timing instead of simple non-blocking loops. 

 

Wokwi Simulator. (2024). TaskScheduler example [Software simulation]. Wokwi. https://wokwi.com/projects/307248583887815233 
 Demonstrates the TaskScheduler library running inside an online simulator. 
Difference: We run equivalent experiments on real Arduino hardware to collect physical timing data. 
