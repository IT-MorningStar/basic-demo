public class CpuDisorder {

    private static int a;
    private static int b;
    private static int x;
    private static int y;

    public static void main(String[] args) {
        int i = 0;
        for (; ; ) {

            a = 5;
            b = 5;
            y = 0;
            x = 0;
            new Thread(() -> {
                a = 12;
                x = b;
            }).start();
            new Thread(() -> {
                b = 12;
                y = a;
            }).start();

            if (x == 5 && y == 5) {
                System.out.printf("循环了%d次\n", i);
                break;
            }
            i++;
        }
    }
}
