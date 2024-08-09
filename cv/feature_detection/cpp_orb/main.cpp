#include <opencv2/core.hpp>
#include <opencv2/imgcodecs.hpp>
#include <opencv2/highgui.hpp>
#include <opencv2/features2d.hpp>

#include <iostream>

#define CAM_ID 0
 
using namespace cv;
 
int main() {
    Mat frame, descriptors, out_frame;

    VideoCapture cap;

    cap.open(CAM_ID, CAP_ANY);

    if (!cap.isOpened()) {
        std::cout << "Unable to open camera!" << std::endl;
        return -1;
    }

    std::vector<KeyPoint> keypoints;

    Ptr<FeatureDetector> detector = ORB::create();
    Ptr<DescriptorExtractor> descriptor = ORB::create();

    while(1) {
        cap.read(frame);
        if (frame.empty()) {
            std::cout << "Blank frame 2!" << std::endl;
            return -1;
        }

        detector->detect(frame, keypoints);

        descriptor->compute(frame, keypoints, descriptors);

        drawKeypoints(frame, keypoints, out_frame);

        imshow("Live", out_frame);
        if (waitKey(5) >= 0) {
            break;
        }
    }

    return 0;
}
