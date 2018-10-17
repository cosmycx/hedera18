/*
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 */

package ui;

import javax.swing.JApplet;

public class screenJApplet extends JApplet {

    /**
     * Initialization method that will be called after the applet is loaded
     * into the browser.
     */
    public void init() {
        // TODO start asynchronous download of heavy resources
        java.awt.EventQueue.invokeLater(new Runnable() {
            public void run() {
            	screenJFrame frame = new screenJFrame();
                frame.setVisible(true);            
            }
        });
    }

    // TODO overwrite start(), stop() and destroy() methods

}
