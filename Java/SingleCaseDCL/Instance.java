public class Instance {
    private static volatile Instance instance;


    private void instance() {
    }
    
    public static Instance getInstance() {
        if (instance == null) {
            // synchronized 加锁放在这，锁的粒度比较细
            synchronized (Instance.class) {
                if (instance == null) {
                    instance = new Instance();
                }
            }
        }
        return instance;
    }

}
