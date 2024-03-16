using System;
using System.Diagnostics;
using System.Threading;

class Program
{
    static void Main()
    {
        Part1();
        Part2();
    }

    static void Part1()
    {
        int[] sizes = { 2, 4, 8, 12, 16, 20, 1000, 5000, 20000 };
        int[] threads = { 1, 2, 4, 8, 12, 16 };

        foreach (int size in sizes)
        {
            foreach (int threadCount in threads)
            {
                Console.WriteLine($"Размерность: {size}, Потоки: {threadCount}");

                int[] A = GenerateArray(size);

                double[] B = new double[A.Length];
                object lockObj = new object(); // объект блокировки для доступа к общему массиву B

                Stopwatch stopwatch = new Stopwatch();
                stopwatch.Start();

                Thread[] workerThreads = new Thread[threadCount];
                for (int i = 0; i < threadCount; i++)
                {
                    int start = i * (A.Length / threadCount);
                    int end = (i == threadCount - 1) ? A.Length : (i + 1) * (A.Length / threadCount);
                    workerThreads[i] = new Thread(() =>
                    {
                        for (int j = start; j < end; j++)
                        {
                            double sum = 0;
                            for (int k = 0; k <= j; k++)
                            {
                                sum += Math.Pow(A[k], 1.789 * j);
                            }
                            lock (lockObj)
                            {
                                B[j] = sum; // Присваиваем результат непосредственно в соответствующий элемент B
                            }
                        }
                    });
                    workerThreads[i].Start();
                }

                foreach (var thread in workerThreads)
                {
                    thread.Join(); // Ожидание завершения всех потоков
                }

                stopwatch.Stop();

                Console.WriteLine($"Время выполнения: {stopwatch.ElapsedMilliseconds} мс");
                Console.WriteLine($"Операций на поток: {A.Length / threadCount}\n");
            }
        }
    }

    static void Part2()
    {
        double a = 0; // Нижний предел интегрирования
        double b = 1; // Верхний предел интегрирования
        double epsilon = 1E-06; // Точность интегрирования
        int maxThreads = 20; // Максимальное число потоков для исследования

        // Интегрирование последовательно
        double resultSequential = CalculateIntegralSequential(a, b, epsilon);
        Console.WriteLine($"Последовательное вычисление: Результат = {resultSequential}, Время выполнения = {GetElapsedTime(() => CalculateIntegralSequential(a, b, epsilon))} мс\n");

        // Исследование зависимости времени работы от числа потоков и точности
        Console.WriteLine("Исследование зависимости времени работы от числа потоков и точности:\n");
        for (int threads = 1; threads <= maxThreads; threads *= 2)
        {
            Console.WriteLine($"Потоки: {threads}");
            double accuracy = 1.0;
            while (true)
            {
                double resultParallel = CalculateIntegralParallel(a, b, accuracy, threads);
                double time = GetElapsedTime(() => CalculateIntegralParallel(a, b, accuracy, threads));
                Console.WriteLine($"Точность: {accuracy}, Результат = {resultParallel}, Время выполнения = {time} мс");
                if (time > 120000) // Ограничение времени выполнения до 2 минут
                {
                    Console.WriteLine("Время выполнения превысило 2 минуты. Остановка исследования.");
                    break;
                }
                if (accuracy <= epsilon)
                    break;
                accuracy /= 10.0;
            }
            Console.WriteLine();
        }
    }

    static double CalculateIntegralSequential(double a, double b, double epsilon)
    {
        int n = 1;
        double integral = 0;
        double prevIntegral = double.MaxValue;
        while (Math.Abs(integral - prevIntegral) > epsilon)
        {
            prevIntegral = integral;
            double h = (b - a) / n;
            integral = 0;
            for (int i = 0; i < n; i++)
            {
                double x = a + i * h;
                integral += Function(x) * h;
            }
            n *= 2;
        }
        return integral;
    }

    static double CalculateIntegralParallel(double a, double b, double epsilon, int threads)
    {
        int n = 1;
        double integral = 0;
        double prevIntegral = double.MaxValue;
        object lockObj = new object();

        while (Math.Abs(integral - prevIntegral) > epsilon)
        {
            prevIntegral = integral;
            double h = (b - a) / n;
            integral = 0;
            Thread[] workerThreads = new Thread[threads];

            for (int i = 0; i < threads; i++)
            {
                int start = i * (n / threads);
                int end = (i == threads - 1) ? n : (i + 1) * (n / threads);
                workerThreads[i] = new Thread(() =>
                {
                    double localIntegral = 0;
                    for (int j = start; j < end; j++)
                    {
                        double x = a + j * h;
                        localIntegral += Function(x) * h;
                    }
                    lock (lockObj)
                    {
                        integral += localIntegral;
                    }
                });
                workerThreads[i].Start();
            }

            foreach (var thread in workerThreads)
            {
                thread.Join(new TimeSpan(0, 0, 2)); // Ограничение времени выполнения каждого потока до 2 секунд
            }

            n *= 2;
        }
        return integral;
    }

    static double Function(double x)
    {
        return x * x; // Пример функции, которую мы интегрируем
    }

    static double GetElapsedTime(Action action)
    {
        Stopwatch stopwatch = new Stopwatch();
        stopwatch.Start();
        action.Invoke();
        stopwatch.Stop();
        return stopwatch.ElapsedMilliseconds;
    }

    static int[] GenerateArray(int size)
    {
        Random rand = new Random();
        int[] arr = new int[size];
        for (int i = 0; i < size; i++)
        {
            arr[i] = rand.Next(1, 100); // Генерация случайных чисел от 1 до 100
        }
        return arr;
    }
}
