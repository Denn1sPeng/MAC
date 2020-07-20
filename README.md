# MAC
    >China Unionpay MAC(Message Authentication Code) generator, include standers about X99/X919/PBOC_DES/PBOC_3DES
    ###中国银联加密机MAC生成
        *在与银行通信过程中，需要使用MAC（Message Authentication Code，报文校验码）验证数据完整性与真实性
        *ANSI标准包括X99/X919，其中x919为x99协议的升级版
        *PBOC（People Bank Of China，中国人民银行）标准包括DES/3DES，3DES为DES的改良版
    ###其他
        *整理这段代码是因当年与银行合作的惨痛经历，银行工作人员对于其加密机制完全没有了解，提供的建议是在机房购买加密机。
        *作者在实践中仅使用过ANSI两个标准，PBOC标准是在整理代码时在网上看到的，也就一并整理添加到一起，但并没有实际应用过
        *在与银行的交互中，秘钥处理也会涉及加解密处理，但因为各银行处理方式不同，在此不做展开
        *整理过程中参考了部分网络代码，感谢各位老铁的开源精神
    ###Links
        *[常见的MAC算法（PBOC_3DES_MAC、ANSI X9.9MAC算法、ANSI x9.19算法）]
(https://blog.csdn.net/kxlele/article/details/84854895?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.nonecase)
        *[algorithm](https://github.com/sunvim/algorithm)